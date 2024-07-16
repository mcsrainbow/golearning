package util_all_done

import (
	"fmt"
	"time"
)

// 模拟运行任务的函数, 传入任务 ID, 返回任务结果
func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)               // 稍作延时, 模拟任务耗时
	return fmt.Sprintf("The result is from %d", id) // 返回包含任务 ID 的结果字符串
}

// 获取第一个完成的任务结果
func FirstResponse() string {
	numOfRunner := 10                    // 定义同时运行的任务数目 10
	ch := make(chan string, numOfRunner) // 创建带缓冲的通道, 用来接收任务结果
	for i := 0; i < numOfRunner; i++ {   // 启动多个 goroutine 运行任务
		go func(i int) {
			ret := runTask(i) // 执行任务
			ch <- ret         // 将结果发送到通道
		}(i)
	}
	return <-ch // 返回第一个完成的任务结果
}

// 获取所有任务的结果
func AllResponse() string {
	numOfRunner := 10                    // 定义同时运行的任务数目 10
	ch := make(chan string, numOfRunner) // 创建带缓冲的通道, 用来接收任务结果
	for i := 0; i < numOfRunner; i++ {   // 启动多个 goroutine 运行任务
		go func(i int) {
			ret := runTask(i) // 执行任务
			ch <- ret         // 将结果发送到通道
		}(i)
	}
	finalRet := ""                     // 用来存储最终结果的字符串
	for j := 0; j < numOfRunner; j++ { // 收集所有任务的结果
		finalRet += <-ch + "\n" // 将每个任务结果追加到最终结果中
	}
	return finalRet // 返回所有任务结果的拼接
}

// 测试获取第一个完成任务
