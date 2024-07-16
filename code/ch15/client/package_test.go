package client

import (
	"ch15/series" // 导入 "ch15/series" 包
	"testing"     // 导入 "testing" 包, 用于编写测试代码
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(5)) // 使用 t.Log 打印 Fibonacci 数列, 参数为 5
	t.Log(series.Square(5))            // 使用 t.Log 打印数字 5 的平方
}
