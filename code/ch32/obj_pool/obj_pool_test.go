package object_pool

import (
	"fmt"
	"testing"
	"time"
)

// 测试对象池功能
func TestObjPool(t *testing.T) {
	pool := NewObjPool(10) // 创建一个容量为 10 的对象池

	// 尝试放置超出池大小的对象
	// if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	// 	t.Error(err)
	// }

	// 循环获取和释放对象
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil { // 从对象池中获取对象, 超时时间为 1 秒
			t.Error(err) // 获取对象失败时记录错误
		} else {
			fmt.Printf("%T\n", v)                      // 打印对象类型
			if err := pool.ReleaseObj(v); err != nil { // 将对象释放回对象池中
				t.Error(err) // 释放对象失败时记录错误
			}
		}
	}

	fmt.Println("Done") // 测试完成
}
