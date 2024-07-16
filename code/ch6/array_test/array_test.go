package array_test

import "testing"

// TestArrayInit: 测试数组初始化
func TestArrayInit(t *testing.T) {
	// 定义一个长度为 3 的 int 类型数组, 默认值为 0
	var arr [3]int
	// 定义一个长度为 4 的 int 类型数组并初始化
	arr1 := [4]int{1, 2, 3, 4}
	// 通过省略符号定义数组并初始化, 根据初始值确定数组长度
	arr3 := [...]int{1, 3, 4, 5}
	// 修改 arr1 的第二个元素的值
	arr1[1] = 5
	// 打印 arr 的第二个和第三个元素的值
	t.Log(arr[1], arr[2])
	// 打印 arr1 和 arr3 数组的值
	t.Log(arr1, arr3)
}

// TestArrayTravel: 测试数组遍历
func TestArrayTravel(t *testing.T) {
	// 定义并初始化数组
	arr3 := [...]int{1, 3, 4, 5}
	// 使用索引遍历数组并打印元素值
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	// 使用 range 关键字遍历数组并打印元素值
	for _, e := range arr3 {
		t.Log(e)
	}
}

// TestArraySection: 测试数组截取
func TestArraySection(t *testing.T) {
	// 定义并初始化数组
	arr3 := [...]int{1, 2, 3, 4, 5}
	// 截取数组的全部元素
	arr3_sec := arr3[:]
	// 打印截取后的数组
	t.Log(arr3_sec)
}
