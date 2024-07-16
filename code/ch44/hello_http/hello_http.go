package main

import (
	"fmt"      // 导入 "fmt" 包, 用于格式化输入输出
	"net/http" // 导入 "net/http" 包, 用于 HTTP 客户端和服务器功能
	"time"     // 导入 "time" 包, 用于时间操作
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // 定义根路径 "/" 的 HTTP 处理函数
		fmt.Fprintf(w, "Hello World!") // 写入 "Hello World!" 到响应
	})
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) { // 定义 "/time/" 路径的 HTTP 处理函数
		t := time.Now()                                 // 获取当前时间
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t) // 将时间格式化为 JSON 字符串
		w.Write([]byte(timeStr))                        // 将字符串写入响应
	})

	http.ListenAndServe(":8080", nil) // 启动 HTTP 服务器, 监听端口 8080
}
