package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

// 定义 Employee 结构体
type Employee struct {
	Id   string // 定义 Id 字段
	Name string // 定义 Name 字段
	Age  int    // 定义 Age 字段
}

// 方法: 定义 Employee 的 String 方法, 方法接收者是值类型
func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))          // 输出 Name 字段的内存地址
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age) // 返回格式化的字符串
}

// 测试函数: 创建 Employee 对象
func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}         // 使用简洁语法创建 Employee 对象
	e1 := Employee{Name: "Mike", Age: 30} // 使用带字段名的方式创建 Employee 对象
	e2 := new(Employee)                   // 返回指针的方式创建 Employee 对象
	e2.Id = "2"                           // 设置 e2 对象的 Id 字段
	e2.Age = 22                           // 设置 e2 对象的 Age 字段
	e2.Name = "Rose"                      // 设置 e2 对象的 Name 字段
	t.Log(e)                              // 打印 e 对象的值
	t.Log(e1)                             // 打印 e1 对象的值
	t.Log(e1.Id)                          // 打印 e1 对象的 Id 字段
	t.Log(e2)                             // 打印 e2 对象的值
	t.Logf("e is %T", e)                  // 打印 e 对象的类型
	t.Logf("e2 is %T", e2)                // 打印 e2 对象的类型
}

// 测试函数: 结构体操作
func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}                          // 创建 Employee 对象
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name)) // 输出 Name 字段的内存地址
	t.Log(e.String())                                      // 打印 e 对象的 String 方法返回的结果
}
