package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

const (
	col = 10000 // 定义常量列数
	row = 10000 // 定义常量行数
)

// 填充矩阵
func fillMatrix(m *[row][col]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano())) // 创建一个新的随机数生成器

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = s.Intn(100000) // 将矩阵的每个元素设置为随机数
		}
	}
}

// 计算矩阵
func calculate(m *[row][col]int) {
	for i := 0; i < row; i++ {
		tmp := 0
		for j := 0; j < col; j++ {
			tmp += m[i][j] // 计算每行的元素和
		}
	}
}

func main() {
	// 创建输出文件
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err) // 不能创建 CPU profile 文件, 打印错误信息并终止程序
	}

	// 获取系统信息
	if err := pprof.StartCPUProfile(f); err != nil { // 监控 CPU
		log.Fatal("could not start CPU profile: ", err) // 无法启动 CPU profile, 打印错误信息并终止程序
	}
	defer pprof.StopCPUProfile() // 在函数退出之前停止 CPU profile

	// 主逻辑区, 进行一些简单的代码运算
	x := [row][col]int{} // 创建一个二维数组
	fillMatrix(&x)       // 填充数组
	calculate(&x)        // 计算数组

	// 创建输出文件
	f1, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err) // 不能创建内存 profile 文件, 打印错误信息并终止程序
	}
	runtime.GC()                                       // 进行垃圾回收, 获取最新的数据信息
	if err := pprof.WriteHeapProfile(f1); err != nil { // 写入内存 profile 信息
		log.Fatal("could not write memory profile: ", err) // 无法写入内存 profile 文件, 打印错误信息并终止程序
	}
	f1.Close() // 关闭文件

	// 创建输出文件
	f2, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("could not create groutine profile: ", err) // 不能创建 goroutine profile 文件, 打印错误信息并终止程序
	}

	if gProf := pprof.Lookup("goroutine"); gProf == nil { // 查找 goroutine profile
		log.Fatal("could not write groutine profile: ") // 没有找到 goroutine profile, 打印错误信息并终止程序
	} else {
		gProf.WriteTo(f2, 0) // 将 goroutine profile 写入文件
	}
	f2.Close() // 关闭文件
}
