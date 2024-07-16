package flexible_reflect

import (
	"errors"
	"reflect"
	"testing"
)

// 测试 DeepEqual 函数
func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	// t.Log(a == b) // 直接比较 map 是无效的
	t.Log(reflect.DeepEqual(a, b)) // 使用 reflect.DeepEqual 比较两个 map

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2)) // 比较两个相同的 slice
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3)) // 比较两个不同顺序的 slice
}

// Employee 结构体定义
type Employee struct {
	EmployeeID string
	Name       string `format:"normal"` // 使用 struct 标签
	Age        int
}

// 更新年龄方法
func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

// Customer 结构体定义
type Customer struct {
	CookieID string
	Name     string
	Age      int
}

// 通过设置值填充结构体
func fillBySettings(st interface{}, settings map[string]interface{}) error {
	// 确保传入的参数是指针类型
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("the first param should be a pointer to the struct type.")
	}
	// Elem() 获取指向的值并确保是结构体
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("the first param should be a pointer to the struct type.")
	}

	// 检查设置是否为 nil
	if settings == nil {
		return errors.New("settings is nil.")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	// 遍历设置 map 并填充结构体
	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue // 如果没有对应字段则跳过
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()                          // 获取指向的值
			vstr.FieldByName(k).Set(reflect.ValueOf(v)) // 设置字段值
		}
	}
	return nil
}

// 测试填充 Name 和 Age 字段
func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 30}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}
