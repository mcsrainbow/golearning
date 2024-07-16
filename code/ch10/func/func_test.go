package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 定义一个返回多个值的函数 `returnMultiValues`
func returnMultiValues() (int, int) {
	// 返回两个随机整数
	return rand.Intn(10), rand.Intn(20)
}

// 定义一个计算函数执行时间的函数 `timeSpent`
func timeSpent(inner func(op int) int) func(op int) int {
	// 返回一个函数, 该函数会计算 `inner` 函数的执行时间
	return func(n int) int {
		start := time.Now()                                     // 记录开始时间
		ret := inner(n)                                         // 执行传入的函数
		fmt.Println("time spent:", time.Since(start).Seconds()) // 输出执行时间
		return ret                                              // 返回函数的执行结果
	}
}

// 定义一个耗时的函数 `slowFun`
func slowFun(op int) int {
	time.Sleep(time.Second * 1) // 睡眠 1 秒
	return op                   // 返回传入的参数
}

// 定义一个测试函数 `TestFn`, 用于测试上述函数
func TestFn(t *testing.T) {
	a, _ := returnMultiValues() // 调用 `returnMultiValues` 函数, 忽略第二个返回值
	t.Log(a)                    // 打印第一个返回值
	tsSF := timeSpent(slowFun)  // 调用 `timeSpent` 包装 `slowFun` 函数
	t.Log(tsSF(10))             // 打印 `tsSF` 函数执行的结果
}

// 定义一个可变参数函数 `Sum`
func Sum(ops ...int) int {
	ret := 0                 // 初始化结果变量
	for _, op := range ops { // 遍历所有参数
		ret += op // 累加参数
	}
	return ret // 返回累加结果
}

// 定义一个测试函数 `TestVarParam`, 用于测试 `Sum` 函数
func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))    // 测试 `Sum` 算和
	t.Log(Sum(1, 2, 3, 4, 5)) // 测试 `Sum` 算和
}

// 定义一个用于清理资源的函数 `Clear`
func Clear() {
	fmt.Println("Clear resources.") // 打印清理信息
}

// 定义一个测试函数 `TestDefer`, 用于测试 defer 语句
func TestDefer(t *testing.T) {
	defer Clear()        // 在函数结束时调用 `Clear` 函数
	fmt.Println("Start") // 打印开始信息
	panic("err")         // 模拟一个 panic 异常
}
