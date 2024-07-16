// Package pipefilter is to define the interfaces and the structures for pipe-filter style implementation
// 包 pipefilter 用于定义管道-过滤器风格实现中的接口和结构

package pipefilter

// Request is the input of the filter
// "Request" 是过滤器的输入

type Request interface{}

// Response is the output of the filter
// "Response" 是过滤器的输出

type Response interface{}

// Filter interface is the definition of the data processing components
// "Filter" 接口定义了数据处理组件

// Pipe-Filter structure
// 管道-过滤器结构

type Filter interface {
	Process(data Request) (Response, error)
}
