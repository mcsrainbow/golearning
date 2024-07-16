package concurrency

import (
	"fmt"
	"testing"
	"time"
)

// 检查取消状态的函数: 如果取消通道有数据, 返回 true , 否则返回 false
func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

// 第一种取消方法: 向取消通道发送一个空结构体
func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

// 第二种取消方法: 关闭取消通道
func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

// 测试取消功能的函数
func TestCancel(t *testing.T) {
	// 创建一个无缓冲的取消通道
	cancelChan := make(chan struct{}, 0)
	// 启动 5 个 goroutine
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				// 检查是否取消
				if isCancelled(cancelCh) {
					break
				}
				// 每次循环睡眠 5 毫秒
				time.Sleep(time.Millisecond * 5)
			}
			// 打印 goroutine 编号和 "Cancelled" 字符串
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	// 调用第二种取消方法, 关闭取消通道
	cancel_2(cancelChan)
	// 主 goroutine 睡眠 1 秒, 以确保所有 goroutine 完成
	time.Sleep(time.Second * 1)
}
