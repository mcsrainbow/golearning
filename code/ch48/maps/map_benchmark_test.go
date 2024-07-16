// package maps 定义当前 package 名称为 "maps"
package maps

import (
	"strconv" // 导入 "strconv", 用于字符串和其他类型之间的转换
	"sync"    // 导入 "sync", 用于并发控制
	"testing" // 导入 "testing", 用于编写测试代码
)

const (
	NumOfReader = 100 // 定义常量 "NumOfReader", 表示读操作的 goroutine 数量
	NumOfWriter = 10  // 定义常量 "NumOfWriter", 表示写操作的 goroutine 数量
)

// 定义接口 "Map", 包含 Set, Get 和 Del 方法, 用于操作 key-value 对
type Map interface {
	Set(key interface{}, val interface{})    // Set 方法, 设置 key 对应的值
	Get(key interface{}) (interface{}, bool) // Get 方法, 获取 key 对应的值, 返回值和是否存在的布尔值
	Del(key interface{})                     // Del 方法, 删除 key 对应的值
}

// 定义函数 "benchmarkMap", 参数 "b *testing.B" 是基准测试对象, "hm Map" 是实现了 Map 接口的实例
func benchmarkMap(b *testing.B, hm Map) {
	for i := 0; i < b.N; i++ { // 遍历 b.N 次, b.N 由测试框架设置
		var wg sync.WaitGroup              // 定义 "wg" 为 WaitGroup, 用于等待所有子 goroutine 结束
		for i := 0; i < NumOfWriter; i++ { // 创建 NumOfWriter 个写操作的 goroutine
			wg.Add(1)   // 增加 WaitGroup 计数器
			go func() { // 启动匿名 goroutine
				for i := 0; i < 100; i++ { // 每个 goroutine 进行 100 次写操作和删除操作
					hm.Set(strconv.Itoa(i), i*i) // 设置 key 和 value
					hm.Set(strconv.Itoa(i), i*i) // 再次设置 key 和 value
					hm.Del(strconv.Itoa(i))      // 删除 key
				}
				wg.Done() // 标记当前 goroutine 完成
			}()
		}
		for i := 0; i < NumOfReader; i++ { // 创建 NumOfReader 个读操作的 goroutine
			wg.Add(1)   // 增加 WaitGroup 计数器
			go func() { // 启动匿名 goroutine
				for i := 0; i < 100; i++ { // 每个 goroutine 进行 100 次读操作
					hm.Get(strconv.Itoa(i)) // 获取 key 对应的值
				}
				wg.Done() // 标记当前 goroutine 完成
			}()
		}
		wg.Wait() // 等待所有的 goroutine 完成
	}
}

// 定义函数 "BenchmarkSyncmap", 用于基准测试不同的 Map 实现
func BenchmarkSyncmap(b *testing.B) {
	b.Run("map with RWLock", func(b *testing.B) { // 定义基准测试子任务 "map with RWLock"
		hm := CreateRWLockMap() // 创建使用 RWLock 的 Map 实现
		benchmarkMap(b, hm)     // 运行基准测试
	})

	b.Run("sync.map", func(b *testing.B) { // 定义基准测试子任务 "sync.map"
		hm := CreateSyncMapBenchmarkAdapter() // 创建 sync.Map 的适配器
		benchmarkMap(b, hm)                   // 运行基准测试
	})

	b.Run("concurrent map", func(b *testing.B) { // 定义基准测试子任务 "concurrent map"
		superman := CreateConcurrentMapBenchmarkAdapter(199) // 创建并发 Map 的适配器
		benchmarkMap(b, superman)                            // 运行基准测试
	})
}
