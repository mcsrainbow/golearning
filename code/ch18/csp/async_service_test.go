package concurrency

import (
	"fmt"
	"testing"
	"time"
)

// 模拟一个服务, 会耗时 50 毫秒
func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

// 模拟另外一个任务, 会输出一些信息并耗时 100 毫秒
func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

// 测试同步服务
func TestService(t *testing.T) {
	fmt.Println(service()) // 调用服务并打印返回值
	otherTask()            // 调用另外一个任务
}

// 异步服务, 返回一个字符串通道
func AsyncService() chan string {
	retCh := make(chan string, 1) // 创建带缓冲区大小为 1 的字符串通道
	go func() {                   // 启动一个匿名 goroutine
		ret := service()                // 调用耗时的服务
		fmt.Println("returned result.") // 打印提示信息
		retCh <- ret                    // 将服务的返回值发送到通道
		fmt.Println("service exited.")  // 打印提示信息
	}()
	return retCh // 返回通道
}

// 测试异步服务
func TestAsynService(t *testing.T) {
	retCh := AsyncService()     // 调用异步服务并获取返回的通道
	otherTask()                 // 调用另外一个任务
	fmt.Println(<-retCh)        // 从通道接收服务的返回值并打印
	time.Sleep(time.Second * 1) // 主 goroutine 休眠 1 秒
}
