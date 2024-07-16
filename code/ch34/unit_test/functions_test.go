package testing

import (
	"fmt"     // 导入 fmt 包, 用于格式化 I/O
	"testing" // 导入 testing 包, 用于编写单元测试

	"github.com/stretchr/testify/assert" // 导入 testify 包中的 assert 库, 用于断言
)

// 测试函数 TestSquare
func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}        // 定义输入数组
	expected := [...]int{1, 4, 9}      // 定义期望的输出数组
	for i := 0; i < len(inputs); i++ { // 循环遍历输入数组
		ret := square(inputs[i]) // 调用 square 函数, 获取返回值
		if ret != expected[i] {  // 比较返回值和期望值
			t.Errorf("input is %d, the expected is %d, the actual %d",
				inputs[i], expected[i], ret) // 如果不一致, 报告错误
		}
	}
}

// 测试函数 TestErrorInCode
func TestErrorInCode(t *testing.T) {
	fmt.Println("Start") // 打印 "Start"
	t.Error("Error")     // 抛出错误, 但继续执行
	fmt.Println("End")   // 打印 "End"
}

// 测试函数 TestFailInCode
func TestFailInCode(t *testing.T) {
	fmt.Println("Start") // 打印 "Start"
	t.Fatal("Error")     // 抛出致命错误, 并中断执行
	fmt.Println("End")   // 此行不会执行
}

// 使用 assert 库的测试函数 TestSquareWithAssert
func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}        // 定义输入数组
	expected := [...]int{1, 4, 9}      // 定义期望的输出数组
	for i := 0; i < len(inputs); i++ { // 循环遍历输入数组
		ret := square(inputs[i])          // 调用 square 函数, 获取返回值
		assert.Equal(t, expected[i], ret) // 使用 assert.Equal 断言返回值与期望值是否相等
	}
}

// 注意: 本文件假定存在一个名为 square 的函数, 其定义不在此文件中
