package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1 // 定义一个 int32 类型的变量 a 并赋值为 1
	var b int64     // 定义一个 int64 类型的变量 b
	b = int64(a)    // 将 int32 类型的变量 a 显式转换为 int64 类型并赋值给 b
	var c MyInt     // 定义一个 MyInt 类型的变量 c
	c = MyInt(b)    // 将 int64 类型的变量 b 显式转换为 MyInt 类型并赋值给 c
	t.Log(a, b, c)  // 输出变量 a, b 和 c 的值
}

func TestPoint(t *testing.T) {
	a := 1     // 定义一个 int 类型的变量 a 并赋值为 1
	aPtr := &a // 定义一个指针 aPtr, 指向变量 a
	//aPtr = aPtr + 1      // 不允许指针运算
	t.Log(a, aPtr)           // 输出变量 a 和 aPtr 的值
	t.Logf("%T %T", a, aPtr) // 输出变量 a 和 aPtr 的类型
}

func TestString(t *testing.T) {
	var s string         // 定义一个字符串类型的变量 s
	t.Log("*" + s + "*") // 字符串 s 的零值是空字符串 ""
	t.Log(len(s))        // 输出字符串 s 的长度
}
