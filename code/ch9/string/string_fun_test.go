package string_test

import (
	"strconv" // 导入 strconv 包, 用于字符串和整数之间的转换
	"strings" // 导入 strings 包, 用于字符串操作
	"testing" // 导入 testing 包, 用于编写测试函数
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"                   // 声明一个字符串 s, 并赋值为 "A,B,C"
	parts := strings.Split(s, ",") // 使用 strings.Split 函数, 以逗号分隔字符串 s
	for _, part := range parts {   // 遍历分隔后的字符串切片 parts
		t.Log(part) // 记录并打印当前遍历到的字符串 part
	}
	t.Log(strings.Join(parts, "-")) // 使用 strings.Join 函数, 以连字符连接字符串切片 parts
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)                         // 将整数 10 转换为字符串, 并赋值给变量 s
	t.Log("str" + s)                              // 记录并打印 "str" 和变量 s 的拼接结果
	if i, err := strconv.Atoi("10"); err == nil { // 将字符串 "10" 转换为整数, 并进行错误检测
		t.Log(10 + i) // 如果无错误, 记录并打印 10 和转换后的整数 i 的和
	}
}
