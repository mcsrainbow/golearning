package empty_interface

import (
	"fmt"
	"testing"
)

// DoSomething 函数接受一个空接口参数 p
func DoSomething(p interface{}) {
	// 使用类型断言的方式检查参数的类型
	// if i, ok := p.(int); ok {
	//  fmt.Println("Integer", i)
	//  return
	// }
	// if s, ok := p.(string); ok {
	//  fmt.Println("stirng", s)
	//  return
	// }
	// fmt.Println("Unknow Type")

	// 使用 type switch 检查参数类型 v 的类型
	switch v := p.(type) {
	case int:
		// 如果类型是 int, 打印 "Integer" 和整数值
		fmt.Println("Integer", v)
	case string:
		// 如果类型是 string, 打印 "String" 和字符串值
		fmt.Println("String", v)
	default:
		// 如果类型不是 int 或 string, 打印 "Unknow Type"
		fmt.Println("Unknow Type")
	}
}

// TestEmptyInterfaceAssertion 是一个测试函数
func TestEmptyInterfaceAssertion(t *testing.T) {
	// 调用 DoSomething 测试整数类型
	DoSomething(10)
	// 调用 DoSomething 测试字符串类型
	DoSomething("10")
}
