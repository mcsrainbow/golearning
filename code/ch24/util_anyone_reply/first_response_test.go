package concurrency

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 模拟执行一个任务并返回结果字符串
func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)               // 模拟任务执行时间
	return fmt.Sprintf("The result is from %d", id) // 返回任务执行结果
}

// 获取第一个任务的返回结果
func FirstResponse() string {
	numOfRunner := 10                    // 设定任务数量
	ch := make(chan string, numOfRunner) // 创建一个缓冲通道
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i) // 执行任务
			ch <- ret         // 将结果发送到通道
		}(i)
	}
	return <-ch // 返回第一个接收到的结果
}

// 测试 FirstResponse 函数
func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine()) // 打印当前的 goroutines 数量
	t.Log(FirstResponse())                   // 打印第一个任务的返回结果
	time.Sleep(time.Second * 1)              // 等待 1 秒
	t.Log("After:", runtime.NumGoroutine())  // 打印等待后的 goroutines 数量
}
