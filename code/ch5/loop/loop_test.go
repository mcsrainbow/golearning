package loop_test

import "testing"

// 定义一个名为 "TestWhileLoop" 的测试函数
func TestWhileLoop(t *testing.T) {
	n := 0
	// 当 n 小于 5 时执行循环
	for n < 5 {
		// 记录当前 n 的值
		t.Log(n)
		// n 增加 1
		n++
	}
}
