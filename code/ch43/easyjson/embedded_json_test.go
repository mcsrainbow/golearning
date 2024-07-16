package jsontest

import (
	"encoding/json" //  导入编码/json包
	"fmt"           //  导入fmt包
	"testing"       //  导入testing包
)

// 定义一个字符串变量, 用来保存要解析的 JSON 字符串
var jsonStr = `{
	"basic_info":{
	  	"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}`

// 定义一个测试函数, 测试标准库的 JSON 编解码
func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)                        //  初始化 Employee 结构体
	err := json.Unmarshal([]byte(jsonStr), e) //  反序列化 JSON 字符串到结构体
	if err != nil {                           //  如果反序列化出错
		t.Error(err) //  打印错误信息
	}
	fmt.Println(*e)                            //  打印反序列化后的结构体内容
	if v, err := json.Marshal(e); err == nil { //  序列化结构体到 JSON 字符串
		fmt.Println(string(v)) //  打印序列化后的 JSON 字符串
	} else {
		t.Error(err) //  如果序列化出错, 打印错误信息
	}
}

// 定义一个测试函数, 测试 EasyJson 的 JSON 编解码
func TestEasyJson(t *testing.T) {
	e := Employee{}                            //  初始化 Employee 结构体
	e.UnmarshalJSON([]byte(jsonStr))           //  使用 EasyJson 反序列化 JSON 字符串
	fmt.Println(e)                             //  打印反序列化后的结构体内容
	if v, err := e.MarshalJSON(); err != nil { //  使用 EasyJson 序列化结构体到 JSON 字符串
		t.Error(err) //  如果序列化出错, 打印错误信息
	} else {
		fmt.Println(string(v)) //  打印序列化后的 JSON 字符串
	}
}

// 定义一个性能基准测试函数, 测试标准库的 JSON 编解码性能
func BenchmarkEmbeddedJson(b *testing.B) {
	b.ResetTimer()             //  重置计时器
	e := new(Employee)         //  初始化 Employee 结构体
	for i := 0; i < b.N; i++ { //  循环 b.N 次
		err := json.Unmarshal([]byte(jsonStr), e) //  反序列化 JSON 字符串到结构体
		if err != nil {                           //  如果反序列化出错
			b.Error(err) //  打印错误信息
		}
		if _, err = json.Marshal(e); err != nil { //  序列化结构体到 JSON 字符串
			b.Error(err) //  如果序列化出错, 打印错误信息
		}
	}
}

// 定义一个性能基准测试函数, 测试 EasyJson 的 JSON 编解码性能
func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()             //  重置计时器
	e := Employee{}            //  初始化 Employee 结构体
	for i := 0; i < b.N; i++ { //  循环 b.N 次
		err := e.UnmarshalJSON([]byte(jsonStr)) //  使用 EasyJson 反序列化 JSON 字符串
		if err != nil {                         //  如果反序列化出错
			b.Error(err) //  打印错误信息
		}
		if _, err = e.MarshalJSON(); err != nil { //  使用 EasyJson 序列化结构体到 JSON 字符串
			b.Error(err) //  如果序列化出错, 打印错误信息
		}
	}
}
