package customer_type

import (
	"fmt"     // 导入格式化库 "fmt"
	"testing" // 导入测试库 "testing"
	"time"    // 导入时间库 "time"
)

// 定义 IntConv 类型, 它是一个接受 int 返回 int 的函数
type IntConv func(op int) int

// 定义时间消耗计算函数 timeSpent, 它接受一个 IntConv 类型的函数并返回同样类型的函数
func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()                                     // 记录开始时间
		ret := inner(n)                                         // 调用传入的函数并获取返回值
		fmt.Println("time spent:", time.Since(start).Seconds()) // 打印消耗的时间
		return ret                                              // 返回传入函数的返回值
	}
}

// 定义一个慢函数 slowFun , 它接受一个 int 参数并返回相同的 int, 且在执行过程中等待 1 秒
func slowFun(op int) int {
	time.Sleep(time.Second * 1) // 等待1秒
	return op                   // 返回参数
}

// 测试函数 TestFn , 用于测试 timeSpent 函数的功能
func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun) // 使用 timeSpent 包装 slowFun 函数
	t.Log(tsSF(10))            // 打印并记录 tsSF(10) 的返回值
}
