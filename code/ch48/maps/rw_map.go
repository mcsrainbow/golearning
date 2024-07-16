package maps

import "sync"

// RWLockMap 结构体, 封装了 map 和读写锁
type RWLockMap struct {
	m    map[interface{}]interface{} // map, 用于存储键值对
	lock sync.RWMutex                // 读写锁, 用于控制并发访问
}

// Get 方法, 获取键对应的值
func (m *RWLockMap) Get(key interface{}) (interface{}, bool) {
	m.lock.RLock()    // 获取读锁
	v, ok := m.m[key] // 从 map 中获取键对应的值
	m.lock.RUnlock()  // 释放读锁
	return v, ok      // 返回值和是否存在标志
}

// Set 方法, 设置键值对
func (m *RWLockMap) Set(key interface{}, value interface{}) {
	m.lock.Lock()    // 获取写锁
	m.m[key] = value // 设置键对应的值
	m.lock.Unlock()  // 释放写锁
}

// Del 方法, 删除键值对
func (m *RWLockMap) Del(key interface{}) {
	m.lock.Lock()    // 获取写锁
	delete(m.m, key) // 删除键对应的值
	m.lock.Unlock()  // 释放写锁
}

// CreateRWLockMap 函数, 创建并返回一个 RWLockMap 实例
func CreateRWLockMap() *RWLockMap {
	m := make(map[interface{}]interface{}, 0) // 创建一个空的 map
	return &RWLockMap{m: m}                   // 返回 RWLockMap 实例指针
}
