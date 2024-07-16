package pipefilter

import (
	"errors"
	"strconv"
)

// 定义一个错误, 表示输入数据应该是 []string 格式
var ToIntFilterWrongFormatError = errors.New("input data should be []string")

// 定义 ToIntFilter 结构体
type ToIntFilter struct {
}

// 创建一个新的 ToIntFilter 实例
func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

// Process 方法用于处理数据, 将字符串数组转换为整数数组
func (tif *ToIntFilter) Process(data Request) (Response, error) {
	// 断言输入数据类型是否为 []string, 如果不是则返回错误
	parts, ok := data.([]string)
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}

	// 创建一个整数数组用于存储转换结果
	ret := []int{}

	// 遍历字符串数组, 将每个字符串转换为整数并添加到结果数组中
	for _, part := range parts {
		s, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}

	// 返回转换后的整数数组
	return ret, nil
}
