package once_test

import (
	"fmt"     // 导入 "fmt" 包, 用于格式化 I/O
	"sync"    // 导入 "sync" 包, 提供同步原语
	"testing" // 导入 "testing" 包, 用于编写测试
	"unsafe"  // 导入 "unsafe" 包, 提供低级别编程功能
)

type Singleton struct {
	data string // 定义 "Singleton" 结构体, 包含一个 "data" 字符串字段
}

var singleInstance *Singleton // 声明一个指向 "Singleton" 的全局指针 "singleInstance"
var once sync.Once            // 声明一个 "sync.Once" 类型的变量 "once"

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")       // 使用一次性操作创建单例对象
		singleInstance = new(Singleton) // 初始化 "singleInstance"
	})
	return singleInstance // 返回已创建的单例对象
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup     // 创建一个 "sync.WaitGroup" 实例 "wg"
	for i := 0; i < 10; i++ { // 循环创建 10 个协程
		wg.Add(1) // 向 "WaitGroup" 中添加一个计数
		go func() {
			obj := GetSingletonObj()                // 调用 "GetSingletonObj" 获取单例对象
			fmt.Printf("%X\n", unsafe.Pointer(obj)) // 打印对象的内存地址
			wg.Done()                               // 完成当前协程的工作, 递减 "WaitGroup" 的计数
		}()
	}
	wg.Wait() // 阻塞当前协程, 直至 "WaitGroup" 的计数器为 0
}
