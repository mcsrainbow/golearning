package main

import (
	"fmt"      // 导入 fmt 包, 用于格式化字符串输出
	"log"      // 导入 log 包, 用于日志记录
	"net/http" // 导入 net/http 包, 用于 HTTP 服务

	"github.com/julienschmidt/httprouter" // 导入第三方 httprouter 包
)

// Index 处理函数: 当访问根路径"/"时, 返回文本"Welcome!\n"
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n") // 向响应写入欢迎文本
}

// Hello 处理函数: 当访问路径"/hello/:name"时, 返回文本"hello, :name!\n"
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name")) // 获取路径参数"name", 并向响应写入问候文本
}

func main() {
	router := httprouter.New() // 创建一个新的 httprouter 路由器实例

	router.GET("/", Index)            // 为根路径"/"注册 Index 处理函数
	router.GET("/hello/:name", Hello) // 为路径"/hello/:name"注册 Hello 处理函数

	log.Fatal(http.ListenAndServe(":8080", router)) // 启动 HTTP 服务器, 监听端口 8080, 并使用 router 处理请求, 如果服务器启动失败, 记录日志错误并退出
}
