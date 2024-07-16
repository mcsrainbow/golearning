package interface_test

import "testing"

// 定义一个 Programmer 接口, 包含一个 WriteHelloWorld 方法
type Programmer interface {
	WriteHelloWorld() string
}

// 定义一个 GoProgrammer 结构体
type GoProgrammer struct {
}

// 为 GoProgrammer 结构体实现 WriteHelloWorld 方法, 返回 "fmt.Println("Hello World")"
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

// 测试函数 TestClient
func TestClient(t *testing.T) {
	// 声明一个 Programmer 接口变量 p
	var p Programmer
	// 将 GoProgrammer 的一个实例赋值给 p
	p = new(GoProgrammer)
	// 调用 p 的 WriteHelloWorld 方法, 打印结果
	t.Log(p.WriteHelloWorld())
}
