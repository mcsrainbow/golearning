package pipefilter

// NewStraightPipeline create a new StraightPipelineWithWallTime
// NewStraightPipeline 创建一个 StraightPipelineWithWallTime 实例
func NewStraightPipeline(name string, filters ...Filter) *StraightPipeline {
	return &StraightPipeline{
		Name:    name,
		Filters: &filters,
	}
}

// StraightPipeline is composed of the filters, and the filters are piled as a straigt line.
// StraightPipeline 由过滤器组成, 过滤器排列成一条直线
type StraightPipeline struct {
	Name    string    // 名称
	Filters *[]Filter // 过滤器数组
}

// Process is to process the coming data by the pipeline
// Process 方法通过管道处理传入的数据
func (f *StraightPipeline) Process(data Request) (Response, error) {
	var ret interface{} // 处理后的返回值
	var err error       // 错误信息
	for _, filter := range *f.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret // 将处理后的数据传递给下一个过滤器
	}
	return ret, err // 返回最终结果和错误信息
}
