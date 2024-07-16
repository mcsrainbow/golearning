package share_mem

import (
	"sync"
	"testing"
	"time"
)

// 测试普通的计数器, 注意: 并不保证线程安全
func TestCounter(t *testing.T) {
	counter := 0 // 初始化计数器
	for i := 0; i < 5000; i++ {
		go func() {
			counter++ // 并发地增加计数器
		}()
	}
	time.Sleep(1 * time.Second)     // 等待 1 秒, 以便所有 goroutine 运行完毕
	t.Logf("counter = %d", counter) // 记录计数器的值
}

// 测试线程安全的计数器, 使用互斥锁确保线程安全
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex // 创建互斥锁
	counter := 0       // 初始化计数器
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock() // 确保函数返回时解锁
			}()
			mut.Lock() // 加锁
			counter++  // 增加计数器
		}()
	}
	time.Sleep(1 * time.Second)     // 等待 1 秒, 以便所有 goroutine 运行完毕
	t.Logf("counter = %d", counter) // 记录计数器的值
}

// 使用 WaitGroup 确保所有 goroutine 完成, 并且计数器线程安全
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex    // 创建互斥锁
	var wg sync.WaitGroup // 创建 WaitGroup
	counter := 0          // 初始化计数器
	for i := 0; i < 5000; i++ {
		wg.Add(1) // 每个 goroutine 增加等待计数
		go func() {
			defer func() {
				mut.Unlock() // 确保函数返回时解锁
			}()
			mut.Lock() // 加锁
			counter++  // 增加计数器
			wg.Done()  // 通知 WaitGroup 该 goroutine 完成
		}()
	}
	wg.Wait()                       // 等待所有 goroutine 完成
	t.Logf("counter = %d", counter) // 记录计数器的值
}
