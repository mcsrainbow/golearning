package profiling

import "testing"

// 测试函数: TestCreateRequest
func TestCreateRequest(t *testing.T) {
	str := createRequest() // 执行 createRequest 函数, 将结果存储到变量 str 中
	t.Log(str)             // 打印 str 变量
}

// 测试函数: TestProcessRequest
func TestProcessRequest(t *testing.T) {
	reqs := []string{}                   // 初始化一个字符串切片变量 reqs
	reqs = append(reqs, createRequest()) // 将 createRequest 函数的结果追加到 reqs 切片
	reps := processRequest(reqs)         // 执行 processRequest 函数, 将结果存储到变量 reps 中
	t.Log(reps[0])                       // 打印 reps 切片中的第一个元素
}

// 基准测试函数: BenchmarkProcessRequest
func BenchmarkProcessRequest(b *testing.B) {
	reqs := []string{}                   // 初始化一个字符串切片变量 reqs
	reqs = append(reqs, createRequest()) // 将 createRequest 函数的结果追加到 reqs 切片
	b.ResetTimer()                       // 重置计时器, 忽略初始化操作的时间
	for i := 0; i < b.N; i++ {           // 循环 b.N 次, 以进行基准测试
		_ = processRequest(reqs) // 执行 processRequest 函数, 忽略返回值
	}
	b.StopTimer() // 停止计时器
}

// 基准测试函数: BenchmarkProcessRequestOld
func BenchmarkProcessRequestOld(b *testing.B) {
	reqs := []string{}                   // 初始化一个字符串切片变量 reqs
	reqs = append(reqs, createRequest()) // 将 createRequest 函数的结果追加到 reqs 切片
	b.ResetTimer()                       // 重置计时器, 忽略初始化操作的时间
	for i := 0; i < b.N; i++ {           // 循环 b.N 次, 以进行基准测试
		_ = processRequestOld(reqs) // 执行 processRequestOld 函数, 忽略返回值
	}
	b.StopTimer() // 停止计时器
}
