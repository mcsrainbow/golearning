package pipefilter

import (
	"errors"
	"strings"
)

// 定义错误信息: "input data should be string"
var SplitFilterWrongFormatError = errors.New("input data should be string")

// SplitFilter 结构体: 保存分隔符
type SplitFilter struct {
	delimiter string
}

// NewSplitFilter 函数: 初始化一个 SplitFilter 实例
func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

// Process 方法: 处理输入数据, 将字符串按照分隔符分割成子字符串切片
func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string) // 检查数据格式/类型, 是否可以处理
	if !ok {
		return nil, SplitFilterWrongFormatError // 返回错误: 如果数据格式不正确
	}
	parts := strings.Split(str, sf.delimiter) // 按照分隔符分割字符串
	return parts, nil
}
