package pipefilter

import "testing"

// 测试 StraightPipeline 功能
func TestStraightPipeline(t *testing.T) {
	// 创建 SplitFilter 过滤器, 使用逗号","作为分隔符
	spliter := NewSplitFilter(",")

	// 创建 ToIntFilter 过滤器, 将字符串转换为整数
	converter := NewToIntFilter()

	// 创建 SumFilter 过滤器, 计算整数数组的和
	sum := NewSumFilter()

	// 创建 StraightPipeline 管道, 包含 spliter, converter 和 sum 过滤器
	sp := NewStraightPipeline("p1", spliter, converter, sum)

	// 处理输入字符串"1,2,3"
	ret, err := sp.Process("1,2,3")

	// 如果处理过程中发生错误, 测试失败并记录错误
	if err != nil {
		t.Fatal(err)
	}

	// 检查处理结果是否等于6, 如果不等于, 测试失败并显示实际结果
	if ret != 6 {
		t.Fatalf("The expected is 6, but the actual is %d", ret)
	}
}
