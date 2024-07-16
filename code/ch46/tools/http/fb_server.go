package main

import (
	"fmt"              // 导入 "fmt" 包, 用于格式化字符串
	"log"              // 导入 "log" 包, 用于记录日志信息
	"net/http"         // 导入 "net/http" 包, 用于 HTTP 服务
	_ "net/http/pprof" // 导入 "net/http/pprof" 包, 启用性能分析工具
)

// GetFibonacciSerie 返回长度为 n 的斐波那契数列
func GetFibonacciSerie(n int) []int {
	ret := make([]int, 2, n) // 创建一个初始长度为 2, 容量为 n 的切片
	ret[0] = 1               // 设置第一个元素为 1
	ret[1] = 1               // 设置第二个元素为 1
	for i := 2; i < n; i++ { // 从索引 2 开始迭代直到 n
		ret = append(ret, ret[i-2]+ret[i-1]) // 计算后续斐波那契数, 并添加到切片
	}
	return ret // 返回生成的斐波那契数列切片
}

// index 处理根路径 "/" 的 HTTP 请求, 返回 "Welcome!" 消息
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!")) // 向客户端发送 "Welcome!" 消息
}

// createFBS 处理 "/fb" 路径的 HTTP 请求, 生成 1,000,000 次长度为 50 的斐波那契数列并返回最后一次结果
func createFBS(w http.ResponseWriter, r *http.Request) {
	var fbs []int                  // 声明一个整型切片, 用于存储斐波那契数列
	for i := 0; i < 1000000; i++ { // 迭代 1,000,000 次
		fbs = GetFibonacciSerie(50) // 每次迭代生成长度为 50 的斐波那契数列
	}
	w.Write([]byte(fmt.Sprintf("%v", fbs))) // 将最后一次生成的斐波那契数列转为字符串并发送给客户端
}

// main 函数是程序的入口
func main() {
	http.HandleFunc("/", index)                  // 设置处理根路径 "/" 请求的处理器为 index 函数
	http.HandleFunc("/fb", createFBS)            // 设置处理 "/fb" 路径请求的处理器为 createFBS 函数
	log.Fatal(http.ListenAndServe(":8081", nil)) // 启动 HTTP 服务器并监听端口 8081, 如果失败则记录错误
}
