package main

import (
	"encoding/json" // 导入用于 JSON 编码和解码的包
	"fmt"           // 导入用于格式化 I/O 的包
	"log"           // 导入用于日志记录的包
	"net/http"      // 导入用于构建 HTTP 客户端和服务端的包

	"github.com/julienschmidt/httprouter" // 导入 httprouter 包, 用于高性能的 HTTP 路由
)

type Employee struct { // 定义 Employee 结构体
	ID   string `json:"id"`   // 员工 ID
	Name string `json:"name"` // 员工姓名
	Age  int    `json:"age"`  // 员工年龄
}

var employeeDB map[string]*Employee // 定义一个全局变量 employeeDB, 用于存储员工信息的映射

func init() { // init 函数, 用于初始化全局变量
	employeeDB = map[string]*Employee{}               // 初始化 employeeDB 为一个空的映射
	employeeDB["Mike"] = &Employee{"e-1", "Mike", 35} // 添加 Mike 的数据到 employeeDB
	employeeDB["Rose"] = &Employee{"e-2", "Rose", 45} // 添加 Rose 的数据到 employeeDB
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) { // 定义 Index 处理函数
	fmt.Fprint(w, "Welcome!\n") // 向客户端响应 "Welcome!\n"
}

func GetEmployeeByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { // 定义 GetEmployeeByName 处理函数
	qName := ps.ByName("name") // 获取 URL 参数中的名字
	var (
		ok       bool      // 定义一个布尔变量, 用于检查员工是否存在
		info     *Employee // 定义一个指向 Employee 的指针变量
		infoJson []byte    // 定义一个字节数组, 用于存储 JSON 编码后的数据
		err      error     // 定义一个 error 类型变量
	)
	if info, ok = employeeDB[qName]; !ok { // 检查员工是否存在
		w.Write([]byte("{\"error\":\"Not Found\"}")) // 如果不存在, 返回错误信息
		return
	}
	if infoJson, err = json.Marshal(info); err != nil { // 将员工信息编码为 JSON 格式
		w.Write([]byte(fmt.Sprintf("{\"error\":,\"%s\"}", err))) // 如果编码出错, 返回错误信息
		return
	}

	w.Write(infoJson) // 将编码后的数据返回给客户端
}

func main() { // 定义主函数
	router := httprouter.New()                       // 创建一个新的 httprouter 实例
	router.GET("/", Index)                           // 注册根路径的处理函数
	router.GET("/employee/:name", GetEmployeeByName) // 注册 /employee/:name 路径的处理函数

	log.Fatal(http.ListenAndServe(":8080", router)) // 启动 HTTP 服务, 监听 8080 端口
}
