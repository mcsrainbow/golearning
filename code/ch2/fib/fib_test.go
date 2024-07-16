package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	// 定义变量: a, 值为 1
	// var a int = 1
	// 定义变量: b, 值为 1
	// var b int = 1
	// 使用分组声明变量: a, 值为 1;b, 值为 1
	// var (
	// 	a int = 1
	// 	b     = 1
	// )
	// 简洁方式声明: a 的值为 1
	a := 1
	// 简洁方式声明: a 的值为 1
	// a := 1
	// 简洁方式声明: b 的值为 1
	b := 1
	// 输出变量 a 的值
	t.Log(a)
	for i := 0; i < 5; i++ {
		// 输出变量 b 的值
		t.Log(" ", b)
		// 使用临时变量 tmp 存储 a 的值
		tmp := a
		// 将变量 b 的值赋给 a
		a = b
		// 将 a 和 tmp 之和赋给 b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	// 简洁方式声明: a 的值为 1
	a := 1
	// 简洁方式声明: b 的值为 2
	b := 2
	// 使用临时变量交换: a, b 的值
	// tmp := a
	// a = b
	// b = tmp
	// 使用多重赋值直接交换: a, b 的值
	a, b = b, a
	// 输出交换后 a 和 b 的值
	t.Log(a, b)
}
