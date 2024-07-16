package map_ext

import "testing"

// TestMapWithFunValue 使用带函数值的 map 的单元测试
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}                 // 创建一个 map, 键类型为 int, 值类型为函数 func(op int) int
	m[1] = func(op int) int { return op }           // 将键 1 对应的值设置为返回输入值的函数
	m[2] = func(op int) int { return op * op }      // 将键 2 对应的值设置为返回输入值平方的函数
	m[3] = func(op int) int { return op * op * op } // 将键 3 对应的值设置为返回输入值立方的函数
	t.Log(m[1](2), m[2](2), m[3](2))                // 打印函数结果: 2, 4, 8
}

// TestMapForSet 使用 map 模拟集合的单元测试
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{} // 创建一个 map, 键类型为 int, 值类型为 bool, 用于模拟集合
	mySet[1] = true         // 将值 1 添加到集合中
	n := 3
	if mySet[n] { // 判断值 3 是否在集合中
		t.Logf("%d is existing", n) // 如果存在, 打印 "3 is existing"
	} else {
		t.Logf("%d is not existing", n) // 如果不存在, 打印 "3 is not existing"
	}
	mySet[3] = true   // 将值 3 添加到集合中
	t.Log(len(mySet)) // 打印集合的长度
	delete(mySet, 1)  // 从集合中删除值 1
	n = 1
	if mySet[n] { // 判断值 1 是否在集合中
		t.Logf("%d is existing", n) // 如果存在, 打印 "1 is existing"
	} else {
		t.Logf("%d is not existing", n) // 如果不存在, 打印 "1 is not existing"
	}
}
