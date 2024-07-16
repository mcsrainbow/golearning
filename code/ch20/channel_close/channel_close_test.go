package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

// dataProducer 函数: 向通道 ch 中写入数据, 并在完成后关闭通道
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ { // 循环写入 0 到 9 的数据
			ch <- i
		}
		close(ch) // 完成写入后关闭通道

		wg.Done() // 通知 WaitGroup 此 goroutine 完成
	}()
}

// dataReceiver 函数: 从通道 ch 中读取数据, 并打印到控制台
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok { // 读取数据并检查通道是否关闭
				fmt.Println(data) // 打印读取到的数据
			} else {
				break // 通道关闭时跳出循环
			}
		}
		wg.Done() // 通知 WaitGroup 此 goroutine 完成
	}()
}

// TestCloseChannel 函数: 测试数据生产者和消费者的协作
func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup // 创建一个 WaitGroup 以等待所有 goroutine 完成
	ch := make(chan int)  // 创建一个整型通道

	wg.Add(1)             // 增加一个计数器, 表示将启动一个生产者 goroutine
	dataProducer(ch, &wg) // 启动生产者 goroutine

	wg.Add(1)             // 增加一个计数器, 表示将启动一个消费者 goroutine
	dataReceiver(ch, &wg) // 启动消费者 goroutine

	// wg.Add(1) // (注释掉) 这行代码是为了增加另一个消费者 goroutine
	// dataReceiver(ch, &wg) // (注释掉) 启动另一个消费者 goroutine

	wg.Wait() // 等待所有 goroutine 完成
}
