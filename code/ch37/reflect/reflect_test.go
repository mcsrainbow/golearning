package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

// 测试获取变量的类型和值
func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f)) // 输出变量的类型和值
	t.Log(reflect.ValueOf(f).Type())             // 输出变量的类型
}

// 检查变量的类型
func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() { // 根据变量的类型种类判断
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

// 测试基本类型
func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(&f) // 检查指向变量的指针的类型
}

// 测试深度比较
func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	//	t.Log(a == b) // map 不能直接比较
	t.Log("a==b?", reflect.DeepEqual(a, b)) // 使用 reflect.DeepEqual 进行比较

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}

	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2)) // 使用 reflect.DeepEqual 进行比较
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3)) // 使用 reflect.DeepEqual 进行比较

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	fmt.Println(c1 == c2)                  // 直接比较结构体
	fmt.Println(reflect.DeepEqual(c1, c2)) // 使用 reflect.DeepEqual 进行比较
}

// 定义 Employee 结构体
type Employee struct {
	EmployeeID string
	Name       string `format:"normal"` // 使用 Tag
	Age        int
}

// 定义 Employee 结构体的方法
func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

// 定义 Customer 结构体
type Customer struct {
	CookieID string
	Name     string
	Age      int
}

// 测试通过名称调用
func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	// 按名字获取成员
	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name")) // 获取 Name 字段的值和类型
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.") // 获取 Name 字段失败
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format")) // 获取 Name 字段的 Tag
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(1)}) // 调用 UpdateAge 方法, 参数为 1
	t.Log("Updated Age:", e) // 输出更新后的 Age
}
