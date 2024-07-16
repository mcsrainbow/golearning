package select_test

import (
	"fmt"
	"testing"
	"time"
)

// service 方法: 模拟一个耗时操作, 延时 500 毫秒
func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

// AsyncService 方法: 启动一个 goroutine 执行耗时操作并返回结果
func AsyncService() chan string {
	retCh := make(chan string, 1) // 创建一个带缓冲的字符串通道, 容量为 1
	//retCh := make(chan string, 1) // 同上, 这行注释掉
	go func() {
		ret := service()                // 调用耗时操作, 获取结果
		fmt.Println("returned result.") // 打印结果返回的提示信息
		retCh <- ret                    // 将结果发送到通道
		fmt.Println("service exited.")  // 打印服务退出的提示信息
	}()
	return retCh // 返回通道
}

// TestSelect 方法: 测试 select 语句与异步操作的组合
func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService(): // 尝试从通道中接收数据
		t.Log(ret) // 如果成功接收到数据, 记录日志
	case <-time.After(time.Millisecond * 100): // 如果 100 毫秒内没有接收到数据
		t.Error("time out") // 记录超时错误
	}
}
