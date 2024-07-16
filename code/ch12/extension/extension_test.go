package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

// 定义 Pet 类型的方法 `Speak`, 输出 "..."
func (p *Pet) Speak() {
	fmt.Print("...")
}

// 定义 Pet 类型的方法 `SpeakTo`, 接受一个字符串参数 host, 先调用 `Speak`, 然后输出 " " 和 host
func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

// 重写 Pet 类型的 `Speak` 方法, 输出 "Wang!"
func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

// 定义测试函数 TestDog
func TestDog(t *testing.T) {
	dog := new(Dog) // 创建一个新的 Dog 对象

	dog.SpeakTo("Chao") // 调用 Dog 对象的 `SpeakTo` 方法, 参数为 "Chao"
}
