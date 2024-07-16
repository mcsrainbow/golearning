package remote

import (
	"testing"

	cm "github.com/easierway/concurrent_map"
)

// 测试函数: TestConcurrentMap uses testing.T
func TestConcurrentMap(t *testing.T) {
	// 创建一个具有 99 个分片的并发 Map: m
	m := cm.CreateConcurrentMap(99)
	// 设置键 "key" 对应的值为 10
	m.Set(cm.StrKey("key"), 10)
	// 记录键 "key" 的值 (应该是 10)
	t.Log(m.Get(cm.StrKey("key")))
}
