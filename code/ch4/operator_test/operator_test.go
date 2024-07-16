package operator_test

import "testing"

const (
	// Readable 代表可读
	Readable = 1 << iota
	// Writable 代表可写
	Writable
	// Executable 代表可执行
	Executable
)

func TestCompareArray(t *testing.T) {
	// 定义数组 a, 内容为 [1, 2, 3, 4]
	a := [...]int{1, 2, 3, 4}
	// 定义数组 b, 内容为 [1, 3, 2, 4]
	b := [...]int{1, 3, 2, 4}
	// 定义数组 c, 内容为 [1, 2, 3, 4, 5] (此行被注释)
	// c := [...]int{1, 2, 3, 4, 5}
	// 定义数组 d, 内容为 [1, 2, 3, 4]
	d := [...]int{1, 2, 3, 4}
	// 打印 a 是否等于 b
	t.Log(a == b)
	// 打印 a 是否等于 c (此行被注释)
	// t.Log(a == c)
	// 打印 a 是否等于 d
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	// 定义变量 a, 初始值为 7 (二进制为 0111)
	a := 7 // 0111
	// 清除 a 中的 Readable (二进制 0010)
	a = a &^ Readable
	// 清除 a 中的 Executable (二进制 0100)
	a = a &^ Executable
	// 打印 a 是否包含 Readable, Writable, Executable 权限
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
