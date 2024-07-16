package lock_test

import (
	"fmt"
	"sync"
	"testing"
)

// 声明一个全局的 map 类型变量 cache
var cache map[string]string

// 声明常量 NUM_OF_READER 和 READ_TIMES
const NUM_OF_READER int = 40
const READ_TIMES = 100000

// 初始化函数, 给 cache 分配内存并赋值
func init() {
	cache = make(map[string]string)

	cache["a"] = "aa"
	cache["b"] = "bb"
}

// 不使用锁进行并发访问
func lockFreeAccess() {

	// 声明 WaitGroup 以等待所有 goroutine 结束
	var wg sync.WaitGroup
	wg.Add(NUM_OF_READER)
	for i := 0; i < NUM_OF_READER; i++ {
		// 启动多个 goroutine 进行并发访问
		go func() {
			for j := 0; j < READ_TIMES; j++ {
				// 读取 cache 中的键 "a" 的值
				_, err := cache["a"]
				if !err {
					fmt.Println("Nothing")
				}
			}
			// 完成一个 goroutine
			wg.Done()
		}()
	}
	// 等待所有的 goroutine 完成
	wg.Wait()
}

// 使用读写锁进行并发访问
func lockAccess() {

	// 声明 WaitGroup 以等待所有 goroutine 结束
	var wg sync.WaitGroup
	wg.Add(NUM_OF_READER)
	// 声明读写锁
	m := new(sync.RWMutex)
	for i := 0; i < NUM_OF_READER; i++ {
		// 启动多个 goroutine 进行并发访问
		go func() {
			for j := 0; j < READ_TIMES; j++ {

				// 上读锁
				m.RLock()
				// 读取 cache 中的键 "a" 的值
				_, err := cache["a"]
				if !err {
					fmt.Println("Nothing")
				}
				// 释放读锁
				m.RUnlock()
			}
			// 完成一个 goroutine
			wg.Done()
		}()
	}
	// 等待所有的 goroutine 完成
	wg.Wait()
}

// 基准测试函数, 测试不使用锁的情况
func BenchmarkLockFree(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lockFreeAccess()
	}
}

// 基准测试函数, 测试使用读写锁的情况
func BenchmarkLock(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lockAccess()
	}
}
