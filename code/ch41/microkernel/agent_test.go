package microkernel

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

// 定义 DemoCollector 结构
type DemoCollector struct {
	// 事件接收器
	evtReceiver EventReceiver

	// 上下文
	agtCtx context.Context

	// 停止通道
	stopChan chan struct{}

	// 名称
	name string

	// 内容
	content string
}

// 新建 DemoCollector 实例
func NewCollect(name string, content string) *DemoCollector {
	return &DemoCollector{
		stopChan: make(chan struct{}),
		name:     name,
		content:  content,
	}
}

// 初始化 DemoCollector
func (c *DemoCollector) Init(evtReceiver EventReceiver) error {
	fmt.Println("initialize collector", c.name) // 打印初始化信息
	c.evtReceiver = evtReceiver
	return nil
}

// 启动 DemoCollector
func (c *DemoCollector) Start(agtCtx context.Context) error {
	fmt.Println("start collector", c.name) // 打印启动信息
	for {
		select {
		case <-agtCtx.Done(): // 当上下文取消时
			c.stopChan <- struct{}{}
			break
		default:
			time.Sleep(time.Millisecond * 50) // 睡眠 50 毫秒
			// 触发事件
			c.evtReceiver.OnEvent(Event{c.name, c.content})
		}
	}
}

// 停止 DemoCollector
func (c *DemoCollector) Stop() error {
	fmt.Println("stop collector", c.name) // 打印停止信息
	select {
	case <-c.stopChan: // 成功停止
		return nil
	case <-time.After(time.Second * 1): // 超时
		return errors.New("failed to stop for timeout") // 返回错误信息
	}
}

// 销毁 DemoCollector
func (c *DemoCollector) Destory() error {
	fmt.Println(c.name, "released resources.")
	return nil
}

// 测试代理
func TestAgent(t *testing.T) {
	agt := NewAgent(100)            // 创建代理实例, 容量为 100
	c1 := NewCollect("c1", "1")     // 新建收集器 c1
	c2 := NewCollect("c2", "2")     // 新建收集器 c2
	agt.RegisterCollector("c1", c1) // 注册 c1 收集器到代理
	agt.RegisterCollector("c2", c2) // 注册 c2 收集器到代理
	if err := agt.Start(); err != nil {
		fmt.Printf("start error %v\n", err) // 若启动代理失败, 打印错误信息
	}
	fmt.Println(agt.Start())    // 打印代理启动信息(注: 可能重复调用)
	time.Sleep(time.Second * 1) // 睡眠 1 秒
	agt.Stop()                  // 停止代理
	agt.Destory()               // 销毁代理
}
