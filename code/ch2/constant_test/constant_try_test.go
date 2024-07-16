package constant_test

import "testing"

// 定义一个常量组, 使用 iota 生成连续的整数常量
const (
	Monday    = 1 + iota // Monday 值为 1 + 0, 即 1
	Tuesday              // Tuesday 值为 1 + 1, 即 2
	Wednesday            // Wednesday 值为 1 + 2, 即 3
)

// 定义另一个常量组, 使用 iota 生成按位左移的常量
const (
	Readable   = 1 << iota // Readable 值为 1 左移 0 位, 即 1
	Writable               // Writable 值为 1 左移 1 位, 即 2
	Executable             // Executable 值为 1 左移 2 位, 即 4
)

// 测试函数, 输出 Monday 和 Tuesday 的值
func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday) // 输出 Monday 和 Tuesday 的值, 分别是 1 和 2
}

// 测试按位与操作的结果是否等于常量
func TestConstantTry1(t *testing.T) {
	a := 1 // 定义变量 a, 值为二进制的 0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	// 输出 a 与三个常量按位与的结果, 分别判断是否等于各自常量的值
	// 结果分别是 true, false, false, 因为 a 的值只有最低位为 1
}
