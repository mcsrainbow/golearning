package main

import (
	"fmt" // 引入 "fmt" 包, 用于格式化 I/O
	"os"  // 引入 "os" 包, 用于操作系统功能
)

func main() {
	if len(os.Args) > 1 { // 如果命令行参数的长度大于 1
		fmt.Println("Hello World", os.Args[1]) // 打印 "Hello World" 和命令行参数的第一个值
	}
}
