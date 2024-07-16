package gc_friendly

import (
	"testing"
)

const NumOfElems = 1000 // 声明常量 "NumOfElems" 为 1000

type Content struct {
	Detail [10000]int // 定义 "Content" 结构体, 包含一个 [10000]int 类型的 "Detail" 字段
}

func withValue(arr [NumOfElems]Content) int {
	// 传递数组 "arr" 的副本
	// fmt.Println(&arr[2]) // 打印数组 "arr" 第 2 个元素的地址
	return 0
}

func withReference(arr *[NumOfElems]Content) int {
	// 传递数组 "arr" 的引用(指针)
	// b := *arr // 取消指针引用, 创建数组 "arr" 的副本
	// fmt.Println(&arr[2]) // 打印数组 "arr" 第 2 个元素的地址
	return 0
}

func TestFn(t *testing.T) {
	var arr [NumOfElems]Content // 声明一个 "NumOfElems" 长度的数组 "arr"
	// fmt.Println(&arr[2]) // 打印数组 "arr" 第 2 个元素的地址
	withValue(arr)      // 以值传递的方式调用 "withValue" 函数
	withReference(&arr) // 以引用传递的方式调用 "withReference" 函数
}

func BenchmarkPassingArrayWithValue(b *testing.B) {
	var arr [NumOfElems]Content // 声明一个 "NumOfElems" 长度的数组 "arr"

	b.ResetTimer() // 重置计时器
	for i := 0; i < b.N; i++ {
		withValue(arr) // 调用 "withValue" 函数
	}
	b.StopTimer() // 停止计时器
}

func BenchmarkPassingArrayWithRef(b *testing.B) {
	var arr [NumOfElems]Content // 声明一个 "NumOfElems" 长度的数组 "arr"

	b.ResetTimer() // 重置计时器
	for i := 0; i < b.N; i++ {
		withReference(&arr) // 调用 "withReference" 函数
	}
	b.StopTimer() // 停止计时器
}
