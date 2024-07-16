package microkernel

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

const (
	Waiting = iota // 定义常量, 表示代理当前的状态为等待
	Running        // 定义常量, 表示代理当前的状态为运行中
)

// 定义错误常量, 当在错误的状态下进行操作时返回
var WrongStateError = errors.New("can not take the operation in the current state")

type CollectorsError struct {
	CollectorErrors []error // 用于存储多个错误的结构体
}

func (ce CollectorsError) Error() string {
	var strs []string // 用于存储错误字符串的切片
	for _, err := range ce.CollectorErrors {
		strs = append(strs, err.Error()) // 将每个错误转换为字符串并添加到切片中
	}
	return strings.Join(strs, ";") // 将多个错误字符串拼接为一个字符串
}

type Event struct {
	Source  string // 事件的来源
	Content string // 事件的内容
}

type EventReceiver interface {
	OnEvent(evt Event) // 定义接口, 要求实现事件处理方法
}

type Collector interface {
	Init(evtReceiver EventReceiver) error // 初始化方法
	Start(agtCtx context.Context) error   // 启动方法
	Stop() error                          // 停止方法
	Destory() error                       // 销毁方法
}

type Agent struct {
	collectors map[string]Collector // 存储采集器的集合
	evtBuf     chan Event           // 事件缓冲通道
	cancel     context.CancelFunc   // 取消函数
	ctx        context.Context      // 上下文对象
	state      int                  // 代理的当前状态
}

func (agt *Agent) EventProcessGroutine() {
	var evtSeg [10]Event // 定义固定大小为10的事件数组
	for {
		for i := 0; i < 10; i++ {
			select {
			case evtSeg[i] = <-agt.evtBuf: // 从事件缓冲通道中获取事件
			case <-agt.ctx.Done(): // 如果上下文被取消, 则退出协程
				return
			}
		}
		fmt.Println(evtSeg) // 打印事件段
	}
}

func NewAgent(sizeEvtBuf int) *Agent {
	agt := Agent{
		collectors: map[string]Collector{},       // 初始化采集器集合
		evtBuf:     make(chan Event, sizeEvtBuf), // 初始化事件缓冲通道
		state:      Waiting,                      // 设置初始状态为等待
	}

	return &agt // 返回新的代理对象
}

func (agt *Agent) RegisterCollector(name string, collector Collector) error {
	if agt.state != Waiting { // 如果代理的状态不是等待
		return WrongStateError // 返回错误
	}
	agt.collectors[name] = collector // 注册采集器
	return collector.Init(agt)       // 初始化采集器
}

func (agt *Agent) startCollectors() error {
	var err error
	var errs CollectorsError // 用于记录所有采集器的错误
	var mutex sync.Mutex     // 定义互斥锁

	for name, collector := range agt.collectors {
		go func(name string, collector Collector, ctx context.Context) {
			defer func() {
				mutex.Unlock() // 释放锁
			}()
			err = collector.Start(ctx) // 启动采集器
			mutex.Lock()               // 加锁
			if err != nil {
				errs.CollectorErrors = append(errs.CollectorErrors, // 记录错误
					errors.New(name+":"+err.Error()))
			}
		}(name, collector, agt.ctx) // 启动协程, 启动采集器
	}
	if len(errs.CollectorErrors) == 0 { // 如果没有任何错误
		return nil
	}
	return errs // 返回所有采集器的错误
}

func (agt *Agent) stopCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		// 停止每个 collector, 如果出错, 将错误信息加入 errs
		if err = collector.Stop(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors,
				errors.New(name+":"+err.Error()))
		}
	}
	if len(errs.CollectorErrors) == 0 {
		// 如果没有错误, 返回 nil
		return nil
	}

	// 返回收集的错误信息
	return errs
}

func (agt *Agent) destoryCollectors() error {
	var err error
	var errs CollectorsError
	for name, collector := range agt.collectors {
		// 销毁每个 collector, 如果出错, 将错误信息加入 errs
		if err = collector.Destory(); err != nil {
			errs.CollectorErrors = append(errs.CollectorErrors,
				errors.New(name+":"+err.Error()))
		}
	}
	if len(errs.CollectorErrors) == 0 {
		// 如果没有错误, 返回 nil
		return nil
	}
	// 返回收集的错误信息
	return errs
}

func (agt *Agent) Start() error {
	if agt.state != Waiting {
		// 如果 Agent 不是处于 Waiting 状态, 返回 WrongStateError
		return WrongStateError
	}
	// 将 Agent 的状态设置为 Running
	agt.state = Running
	// 创建一个新的上下文和取消函数
	agt.ctx, agt.cancel = context.WithCancel(context.Background())
	// 启动事件处理协程
	go agt.EventProcessGroutine()
	// 启动所有的 collector
	return agt.startCollectors()
}

func (agt *Agent) Stop() error {
	if agt.state != Running {
		// 如果 Agent 不是处于 Running 状态, 返回 WrongStateError
		return WrongStateError
	}
	// 将 Agent 的状态设置为 Waiting
	agt.state = Waiting
	// 取消上下文
	agt.cancel()
	// 停止所有的 collector
	return agt.stopCollectors()
}

func (agt *Agent) Destory() error {
	if agt.state != Waiting {
		// 如果 Agent 不是处于 Waiting 状态, 返回 WrongStateError
		return WrongStateError
	}
	// 销毁所有的 collector
	return agt.destoryCollectors()
}

func (agt *Agent) OnEvent(evt Event) {
	// 将事件传递到事件缓冲区
	agt.evtBuf <- evt
}
