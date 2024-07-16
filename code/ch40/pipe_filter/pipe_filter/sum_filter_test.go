package pipefilter

import "testing"

// 定义单元测试函数 TestSumElems
func TestSumElems(t *testing.T) {
	// 创建一个新的 SumFilter 实例
	sf := NewSumFilter()
	// 调用 SumFilter 的 Process 方法并传入整数切片进行处理
	ret, err := sf.Process([]int{1, 2, 3})
	// 如果 Process 方法返回错误, 则终止测试并报告错误
	if err != nil {
		t.Fatal(err)
	}
	// 期望的结果是 6
	expected := 6
	// 如果实际结果不符合期望, 则终止测试并报告错误
	if ret != expected {
		t.Fatalf("The expected is %d, but actual is %d", expected, ret)
	}
}

// 定义单元测试函数 TestWrongInputForSumFilter, 测试传入不正确的输入类型时的行为
func TestWrongInputForSumFilter(t *testing.T) {
	// 创建一个新的 SumFilter 实例
	sf := NewSumFilter()
	// 尝试传入 float32 类型的切片, SumFilter 预计无法处理这种类型的数据
	_, err := sf.Process([]float32{1.1, 2.2, 3.1})

	// 如果 Process 方法没有返回错误, 则终止测试并报告错误
	if err == nil {
		t.Fatal("An error is expected.")
	}
}
