package object_pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// 测试单个使用 sync.Pool
func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int) // 从 pool 获取对象
	fmt.Println(v)
	pool.Put(3)               // 将对象放回 pool
	runtime.GC()              // 触发 GC, GC 会清除 sync.pool 中缓存的对象
	v1, _ := pool.Get().(int) // 再次从 pool 获取对象
	fmt.Println(v1)
}

// 测试多协程下使用 sync.Pool
func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}

	pool.Put(100) // 将对象放入 pool
	pool.Put(100) // 将对象放入 pool
	pool.Put(100) // 将对象放入 pool

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ { // 启动 10 个协程
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get()) // 从 pool 获取对象
			wg.Done()
		}(i)
	}
	wg.Wait() // 等待所有协程完成
}
