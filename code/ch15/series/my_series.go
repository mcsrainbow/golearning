package series

import "fmt"

// func init() { 初始化函数 } // 打开包时, 自动执行
func init() {
	fmt.Println("init1")
}

// func init() { 第二个初始化函数 } // 继续自动执行
func init() {
	fmt.Println("init2")
}

// func Square(n int) int { 计算平方函数 }
func Square(n int) int {
	return n * n
}

// func GetFibonacciSerie(n int) []int { 获取斐波那契数列函数 }
func GetFibonacciSerie(n int) []int {
	// ret := []int{1, 1} // 初始化包含两个元素 "1" 的切片
	ret := []int{1, 1}
	// for i := 2; i < n; i++ { // 遍历循环, 生成斐波那契数列
	for i := 2; i < n; i++ {
		// ret = append(ret; ret[i-2]+ret[i-1]) // 追加新元素到切片: 前两个元素之和
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	// return ret // 返回斐波那契数列
	return ret
}
