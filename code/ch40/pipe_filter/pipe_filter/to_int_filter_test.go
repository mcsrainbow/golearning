package pipefilter

import (
	"reflect" // 导入 reflect 包, 用于深度比较
	"testing" // 导入 testing 包, 用于编写测试用例
)

// 测试函数 TestConvertToInt, 用于测试转换为整数的功能
func TestConvertToInt(t *testing.T) {
	tif := NewToIntFilter()                           // 创建一个新的 ToIntFilter
	resp, err := tif.Process([]string{"1", "2", "3"}) // 处理字符串数组并转换为整数
	if err != nil {
		t.Fatal(err) // 如果发生错误, 打印错误并终止测试
	}
	expected := []int{1, 2, 3}              // 期望得到的结果
	if !reflect.DeepEqual(expected, resp) { // 使用反射包中的 DeepEqual 函数比较结果
		t.Fatalf("The expected is %v, the actual is %v", expected, resp) // 如果结果不一致, 打印期望和实际结果并终止测试
	}
}

// 测试函数 TestWrongInputForTIF, 用于测试异常输入的处理
func TestWrongInputForTIF(t *testing.T) {
	tif := NewToIntFilter()                          // 创建一个新的 ToIntFilter
	_, err := tif.Process([]string{"1", "2.2", "3"}) // 处理包含无效输入的字符串数组
	if err == nil {
		t.Fatal("An error is expected for wrong input") // 如果没有错误, 打印提示并终止测试
	}
	_, err = tif.Process([]int{1, 2, 3}) // 处理错误类型的输入
	if err == nil {
		t.Fatal("An error is expected for wrong input") // 如果没有错误, 打印提示并终止测试
	}
}
