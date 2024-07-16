package pipefilter

import (
	"reflect" // 导入 reflect 包
	"testing" // 导入 testing 包
)

func TestStringSplit(t *testing.T) {
	sf := NewSplitFilter(",")        // 创建一个使用逗号分隔符的 SplitFilter
	resp, err := sf.Process("1,2,3") // 处理字符串 "1,2,3"
	if err != nil {                  // 如果出现错误
		t.Fatal(err) // 记录致命错误并终止测试
	}
	parts, ok := resp.([]string) // 将响应转换为字符串数组
	if !ok {                     // 如果类型转换失败
		t.Fatalf("Repsonse type is %T, but the expected type is string", parts) // 记录致命错误并终止测试: 类型不匹配的错误信息
	}
	if !reflect.DeepEqual(parts, []string{"1", "2", "3"}) { // 比较结果数组是否为 {"1", "2", "3"}
		t.Errorf("Expected value is {\"1\",\"2\",\"3\"}, but actual is %v", parts) // 如果不匹配, 记录错误信息
	}
}

func TestWrongInput(t *testing.T) {
	sf := NewSplitFilter(",") // 创建一个使用逗号分隔符的 SplitFilter
	_, err := sf.Process(123) // 尝试处理一个整数
	if err == nil {           // 如果没有返回错误
		t.Fatal("An error is expected.") // 记录致命错误并终止测试: 预期错误但未发生
	}
}
