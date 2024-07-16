package maps

import "sync"

// 创建一个名为 "CreateSyncMapBenchmarkAdapter" 的函数, 返回 *SyncMapBenchmarkAdapter 的指针
func CreateSyncMapBenchmarkAdapter() *SyncMapBenchmarkAdapter {
	return &SyncMapBenchmarkAdapter{}
}

// 定义一个名为 "SyncMapBenchmarkAdapter" 的结构体, 包含一个 sync.Map 类型的字段 m
type SyncMapBenchmarkAdapter struct {
	m sync.Map
}

// "Set" 方法, 接受 key 和 val 两个参数, 使用 sync.Map 的 Store 方法将键值存储到 map 中
func (m *SyncMapBenchmarkAdapter) Set(key interface{}, val interface{}) {
	m.m.Store(key, val)
}

// "Get" 方法, 接受一个 key 参数, 使用 sync.Map 的 Load 方法从 map 中获取对应的值, 并返回值和是否存在的布尔值
func (m *SyncMapBenchmarkAdapter) Get(key interface{}) (interface{}, bool) {
	return m.m.Load(key)
}

// "Del" 方法, 接受一个 key 参数, 使用 sync.Map 的 Delete 方法从 map 中删除该键
func (m *SyncMapBenchmarkAdapter) Del(key interface{}) {
	m.m.Delete(key)
}
