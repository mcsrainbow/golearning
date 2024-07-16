package concat_string

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const numbers = 100 // 定义常量 "numbers", 值为 100

// 基准测试方法: 使用 fmt.Sprintf 进行字符串连接
func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()                   // 重置计时器
	for idx := 0; idx < b.N; idx++ { // 循环次数为 b.N, 基准测试的默认值
		var s string                   // 声明字符串变量 "s"
		for i := 0; i < numbers; i++ { // 循环 "numbers" 次
			s = fmt.Sprintf("%v%v", s, i) // 使用 fmt.Sprintf 进行字符串连接
		}
	}
	b.StopTimer() // 停止计时器
}

// 基准测试方法: 使用 strings.Builder 进行字符串连接
func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()                   // 重置计时器
	for idx := 0; idx < b.N; idx++ { // 循环次数为 b.N, 基准测试的默认值
		var builder strings.Builder    // 声明 strings.Builder 变量 "builder"
		for i := 0; i < numbers; i++ { // 循环 "numbers" 次
			builder.WriteString(strconv.Itoa(i)) // 将整数转换为字符串并写入 builder

		}
		_ = builder.String() // 将 builder 转为字符串, 不使用结果
	}
	b.StopTimer() // 停止计时器
}

// 基准测试方法: 使用 bytes.Buffer 进行字符串连接
func BenchmarkBytesBuf(b *testing.B) {
	b.ResetTimer()                   // 重置计时器
	for idx := 0; idx < b.N; idx++ { // 循环次数为 b.N, 基准测试的默认值
		var buf bytes.Buffer           // 声明 bytes.Buffer 变量 "buf"
		for i := 0; i < numbers; i++ { // 循环 "numbers" 次
			buf.WriteString(strconv.Itoa(i)) // 将整数转换为字符串并写入 buf
		}
		_ = buf.String() // 将 buf 转为字符串, 不使用结果
	}
	b.StopTimer() // 停止计时器
}

// 基准测试方法: 使用字符串相加进行字符串连接
func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()                   // 重置计时器
	for idx := 0; idx < b.N; idx++ { // 循环次数为 b.N, 基准测试的默认值
		var s string                   // 声明字符串变量 "s"
		for i := 0; i < numbers; i++ { // 循环 "numbers" 次
			s += strconv.Itoa(i) // 将整数转换为字符串并累加到 "s"
		}

	}
	b.StopTimer() // 停止计时器
}
