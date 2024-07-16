package slice_test

import "testing"

// 测试切片初始化
func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0)) // 打印 s0 的长度和容量
	s0 = append(s0, 1)      // 向 s0 追加一个元素 1
	t.Log(len(s0), cap(s0)) // 打印 s0 添加元素后的长度和容量

	s1 := []int{1, 2, 3, 4} // 初始化切片 s1
	t.Log(len(s1), cap(s1)) // 打印 s1 的长度和容量

	s2 := make([]int, 3, 5)           // 使用 make 初始化切片 s2, 长度为 3, 容量为 5
	t.Log(len(s2), cap(s2))           // 打印 s2 的长度和容量
	t.Log(s2[0], s2[1], s2[2])        // 打印 s2 的前三个元素
	s2 = append(s2, 1)                // 向 s2 追加一个元素 1
	t.Log(s2[0], s2[1], s2[2], s2[3]) // 打印 s2 的前四个元素
	t.Log(len(s2), cap(s2))           // 打印 s2 添加元素后的长度和容量
}

// 测试切片的自动增长
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)      // 向 s 中追加元素 i
		t.Log(len(s), cap(s)) // 打印 s 的长度和容量
	}
}

// 测试切片共享底层数组的特性
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"} // 初始化一个包含 12 个月份的切片
	Q2 := year[3:6]                         // Q2 包含 "Apr", "May", "Jun"
	t.Log(Q2, len(Q2), cap(Q2))             // 打印 Q2 的内容, 长度和容量
	summer := year[5:8]                     // summer 包含 "Jun", "Jul", "Aug"
	t.Log(summer, len(summer), cap(summer)) // 打印 summer 的内容, 长度和容量
	summer[0] = "Unknow"                    // 修改 summer 的第一个元素
	t.Log(Q2)                               // 打印 Q2 以查看其第一个元素是否被修改
	t.Log(year)                             // 打印 year 查看变化
}

// 测试切片比较
func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4} // 初始化切片 a
	b := []int{1, 2, 3, 4} // 初始化切片 b
	// if a == b { // 切片只能和 nil 比较
	// 	t.Log("equal")
	// }
	t.Log(a, b) // 打印 a 和 b 的内容
}
