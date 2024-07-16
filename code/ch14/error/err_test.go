package err_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

// 定义一个 错误 变量:  当输入值 小于 2 时 返回
var LessThanTwoError = errors.New("n should be not less than 2")

// 定义一个 错误 变量:  当输入值 大于 100 时 返回
var LargerThenHundredError = errors.New("n should be not larger than 100")

// 定义一个 函数:  获取 斐波那契 数列
func GetFibonacci(n int) ([]int, error) {
	// 如果 输入值 小于 2, 返回 错误 LessThanTwoError
	if n < 2 {
		return nil, LessThanTwoError
	}
	// 如果 输入值 大于 100, 返回 错误 LargerThenHundredError
	if n > 100 {
		return nil, LargerThenHundredError
	}
	// 初始化 斐波那契 数列
	fibList := []int{1, 1}

	// 计算 斐波那契 数列
	for i := 2; /*短变量声明 := */ i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

// 定义一个 测试函数:  测试 GetFibonacci 函数
func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(1); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less.")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

// 定义一个 函数:  获取 斐波那契 数列, 从 字符串 输入
func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)
	// 将 字符串 转换 为 整数
	if i, err = strconv.Atoi(str); err == nil {
		// 计算 斐波那契 数列
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

// 定义一个 函数:  获取 斐波那契 数列, 从 字符串 输入
func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	// 将 字符串 转换 为 整数, 如果 出错, 返回 错误
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	// 计算 斐波那契 数列, 如果 出错, 返回 错误
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}
