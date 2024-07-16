package maps

import "github.com/easierway/concurrent_map"

// ConcurrentMapBenchmarkAdapter 结构体定义, 适配 concurrent_map.ConcurrentMap
type ConcurrentMapBenchmarkAdapter struct {
	cm *concurrent_map.ConcurrentMap // 内部使用的并发映射
}

// Set 方法: 设置键值对到并发映射中
func (m *ConcurrentMapBenchmarkAdapter) Set(key interface{}, value interface{}) {
	// 将 key 转换为字符串类型, 再设置到并发映射中
	m.cm.Set(concurrent_map.StrKey(key.(string)), value)
}

// Get 方法: 从并发映射中获取键对应的值
func (m *ConcurrentMapBenchmarkAdapter) Get(key interface{}) (interface{}, bool) {
	// 将 key 转换为字符串类型, 再获取对应的值
	return m.cm.Get(concurrent_map.StrKey(key.(string)))
}

// Del 方法: 从并发映射中删除键值对
func (m *ConcurrentMapBenchmarkAdapter) Del(key interface{}) {
	// 将 key 转换为字符串类型, 再删除对应的值
	m.cm.Del(concurrent_map.StrKey(key.(string)))
}

// CreateConcurrentMapBenchmarkAdapter 函数: 创建并返回一个 ConcurrentMapBenchmarkAdapter 实例
func CreateConcurrentMapBenchmarkAdapter(numOfPartitions int) *ConcurrentMapBenchmarkAdapter {
	// 创建分区数为 numOfPartitions 的并发映射
	conMap := concurrent_map.CreateConcurrentMap(numOfPartitions)
	// 返回包含并发映射的适配器实例
	return &ConcurrentMapBenchmarkAdapter{conMap}
}
