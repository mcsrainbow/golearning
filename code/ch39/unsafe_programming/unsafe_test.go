package unsafe_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type Customer struct {
	Name string // 客户姓名
	Age  int    // 客户年龄
}

func TestUnsafe(t *testing.T) {
	i := 10
	// 将整型 i 转换为浮点型
	f := *(*float64)(unsafe.Pointer(&i))
	// 输出变量 i 的指针和值
	t.Log(unsafe.Pointer(&i))
	t.Log(f)
}

// The cases is suitable for unsafe
type MyInt int

// 合理的类型转换
func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	// 将 []int 数组转换为 []MyInt 类型
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b)
}

// 原子类型操作
func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	// 写数据的函数
	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		// 原子操作存储指针
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}
	// 读数据的函数
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		// 打印指针和指针指向的数组
		fmt.Println(data, *(*[]int)(data))
	}
	var wg sync.WaitGroup
	writeDataFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
