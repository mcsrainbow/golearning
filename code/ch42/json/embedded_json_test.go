package jsontest

import (
	"encoding/json" // 导入 encoding/json 包, 用于处理 JSON 数据
	"fmt"           // 导入 fmt 包, 用于格式化 I/O
	"testing"       // 导入 testing 包, 用于编写测试
)

// 定义一个 JSON 字符串, 表示员工的基本信息以及工作技能
var jsonStr = `{
	"basic_info":{
	  	"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}	`

// 测试函数, 用于测试嵌套 JSON 的反序列化和序列化
func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)                        // 创建一个新的 Employee 对象
	err := json.Unmarshal([]byte(jsonStr), e) // 将 JSON 字符串反序列化为 Employee 对象
	if err != nil {                           // 如果反序列化出错
		t.Error(err) // 输出错误信息
	}
	fmt.Println(*e)                            // 打印反序列化后的 Employee 对象
	if v, err := json.Marshal(e); err == nil { // 如果成功将 Employee 对象序列化为 JSON
		fmt.Println(string(v)) // 打印序列化后的 JSON 字符串
	} else { // 如果序列化出错
		t.Error(err) // 输出错误信息
	}
}
