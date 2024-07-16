package panic_recover

import (
	"errors"  // 导入 errors 包 , 用于创建错误对象
	"fmt"     // 导入 fmt 包 , 用于格式化 I/O
	"testing" // 导入 testing 包 , 用于编写单元测试
)

func TestPanicVxExit(t *testing.T) {

	// 延迟执行一个匿名函数 , 用于恢复 panic
	defer func() {
		if err := recover(); err != nil { // 如果捕获到 panic , 则打印恢复信息
			fmt.Println("recovered from", err)
		}
	}()
	fmt.Println("Start")                  // 打印 "Start"
	panic(errors.New("Something wrong!")) // 引发一个含有 "Something wrong!" 信息的 panic

	// os.Exit(-1) // 该代码没有执行意义 , 因为前面已经 panic , 可以删除

	// fmt.Println("End") // 该代码没有执行意义 , 因为前面已经 panic , 可以删除
}
