package benchmark_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 定义一个测试函数, 用于测试通过加号连接字符串
func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)                    // 创建一个断言对象
	elems := []string{"1", "2", "3", "4", "5"} // 定义一个字符串切片
	ret := ""                                  // 初始化一个空字符串
	for _, elem := range elems {               // 遍历切片中的每一个元素
		ret += elem // 将元素连接到 ret 字符串
	}
	assert.Equal("12345", ret) // 断言连接后的字符串是否与 "12345" 相等
}

// 定义一个测试函数, 用于测试通过 bytes.Buffer 连接字符串
func TestConcatStringByBytesBuffer(t *testing.T) {
	assert := assert.New(t)                    // 创建一个断言对象
	var buf bytes.Buffer                       // 创建一个 bytes.Buffer 对象
	elems := []string{"1", "2", "3", "4", "5"} // 定义一个字符串切片
	for _, elem := range elems {               // 遍历切片中的每一个元素
		buf.WriteString(elem) // 将元素写入到 buf 缓冲区

	}
	assert.Equal("12345", buf.String()) // 断言连接后的字符串是否与 "12345" 相等
}

// 定义一个基准测试函数, 用于测试通过加号连接字符串的性能
func BenchmarkConcatStringByAdd(b *testing.B) {

	elems := []string{"1", "2", "3", "4", "5"} // 定义一个字符串切片
	b.ResetTimer()                             // 重置计时器
	for i := 0; i < b.N; i++ {                 // 循环进行基准测试
		ret := ""                    // 初始化一个空字符串
		for _, elem := range elems { // 遍历切片中的每一个元素
			ret += elem // 将元素连接到 ret 字符串
		}
	}
	b.StopTimer() // 停止计时器
}

// 定义一个基准测试函数, 用于测试通过 bytes.Buffer 连接字符串的性能
func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"} // 定义一个字符串切片
	b.ResetTimer()                             // 重置计时器
	for i := 0; i < b.N; i++ {                 // 循环进行基准测试
		var buf bytes.Buffer // 创建一个 bytes.Buffer 对象

		for _, elem := range elems { // 遍历切片中的每一个元素
			buf.WriteString(elem) // 将元素写入到 buf 缓冲区

		}
	}
	b.StopTimer() // 停止计时器

}
