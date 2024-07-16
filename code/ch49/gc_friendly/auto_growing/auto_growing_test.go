package gc_friendly

import "testing"

// 常量定义: 每个切片包含的元素数量
const numOfElems = 100000

// 常量定义: 测试重复运行的次数
const times = 1000

func TestAutoGrow(t *testing.T) {
	// 测试: 自动增长切片
	for i := 0; i < times; i++ {
		s := []int{} // 初始化一个空的切片
		for j := 0; j < numOfElems; j++ {
			s = append(s, j) // 向切片中追加元素
		}
	}
}

func TestProperInit(t *testing.T) {
	// 测试: 预分配切片容量
	for i := 0; i < times; i++ {
		s := make([]int, 0, numOfElems) // 初始化一个容量预分配的切片
		for j := 0; j < numOfElems; j++ {
			s = append(s, j) // 向切片中追加元素
		}
	}
}

func TestOverSizeInit(t *testing.T) {
	// 测试: 过度预分配切片容量
	for i := 0; i < times; i++ {
		s := make([]int, 0, 800000) // 初始化一个容量过度预分配的切片
		for j := 0; j < numOfElems; j++ {
			s = append(s, j) // 向切片中追加元素
		}
	}
}

func BenchmarkAutoGrow(b *testing.B) {
	// 基准测试: 自动增长切片
	for i := 0; i < b.N; i++ {
		s := []int{} // 初始化一个空的切片
		for j := 0; j < numOfElems; j++ {
			s = append(s, j) // 向切片中追加元素
		}
	}
}

func BenchmarkProperInit(b *testing.B) {
	// 基准测试: 预分配切片容量
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, numOfElems) // 初始化一个容量预分配的切片
		for j := 0; j < numOfElems; j++ {
			s = append(s, j) // 向切片中追加元素
		}
	}
}

func BenchmarkOverSizeInit(b *testing.B) {
	// 基准测试: 过度预分配切片容量
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, numOfElems*8) // 初始化一个容量过度预分配的切片
		for j := 0; j < numOfElems; j++ {
			s = append(s, j) // 向切片中追加元素
		}
	}
}
