package polymorphism

import (
	"fmt"
	"testing"
)

// 定义一个 Code 类型, 实际是 string 类型的别名
type Code string

// 定义一个 Programmer 接口, 包含一个 WriteHelloWorld() 方法, 返回值类型是 Code
type Programmer interface {
	WriteHelloWorld() Code
}

// 定义一个 GoProgrammer 结构体
type GoProgrammer struct {
}

// 给 GoProgrammer 实现 WriteHelloWorld() 方法, 实现 Programmer 接口
func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")" // 返回 Go 语言的打印 "Hello World!" 代码
}

// 定义一个 JavaProgrammer 结构体
type JavaProgrammer struct {
}

// 给 JavaProgrammer 实现 WriteHelloWorld() 方法, 实现 Programmer 接口
func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")" // 返回 Java 语言的打印 "Hello World!" 代码
}

// 定义一个通用函数, 参数是 Programmer 接口类型, 打印其类型和调用 WriteHelloWorld() 方法的结果
func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

// 测试函数, 测试多态
func TestPolymorphism(t *testing.T) {
	goProg := &GoProgrammer{}       // 创建 GoProgrammer 类型的结构体实例, 并取地址
	javaProg := new(JavaProgrammer) // 使用 new 创建 JavaProgrammer 类型的结构体实例
	writeFirstProgram(goProg)       // 调用 writeFirstProgram(), 传入 goProg 实例
	writeFirstProgram(javaProg)     // 调用 writeFirstProgram(), 传入 javaProg 实例
}
