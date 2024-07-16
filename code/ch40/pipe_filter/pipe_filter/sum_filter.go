package pipefilter

import "errors"

// 定义错误, 表示输入数据格式错误
var SumFilterWrongFormatError = errors.New("input data should be []int")

// SumFilter 结构体定义
type SumFilter struct {
}

// NewSumFilter 函数, 返回一个新的 SumFilter 实例
func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

// Process 方法, 处理输入数据并返回结果
// 参数 "data" 的类型是 Request, 返回值是 Response 和 错误
func (sf *SumFilter) Process(data Request) (Response, error) {
	// 将输入数据转换成 []int 类型
	elems, ok := data.([]int)
	// 检查数据类型, 如果不是 []int, 返回错误
	if !ok {
		return nil, SumFilterWrongFormatError
	}
	// 初始化结果为 0
	ret := 0
	// 循环累加每个元素的值
	for _, elem := range elems {
		ret += elem
	}
	// 返回累加结果
	return ret, nil
}
