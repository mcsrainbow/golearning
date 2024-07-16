package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

// 引入必要的包

func TestGroutine(t *testing.T) {
	// 定义一个名为 TestGroutine 的测试函数, 接收一个 *testing.T 类型的参数

	for i := 0; i < 10; i++ {
		// 循环 10 次, 从 0 到 9

		go func(i int) {
			// 启动一个新的协程, 传入当前的循环变量 i 作为参数
			//time.Sleep(time.Second * 1)
			// 将这一行解注释可以引入 1 秒的延迟

			fmt.Println(i)
			// 打印当前协程中的 i 的值
		}(i)
		// 调用协程函数, 并传入当前的 i
	}
	time.Sleep(time.Millisecond * 50)
	// 主协程休眠 50 毫秒, 以便所有子协程有机会运行
}
