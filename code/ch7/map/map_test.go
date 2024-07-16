package my_map

import "testing"

// TestInitMap 初始化并测试不同类型的 map
func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9} // 定义并初始化一个 map
	t.Log(m1[2])                        // 打印 key 为 2 的值
	t.Logf("len m1=%d", len(m1))        // 打印 m1 的长度
	m2 := map[int]int{}                 // 定义一个空的 map
	m2[4] = 16                          // 设置 key 为 4 的值为 16
	t.Logf("len m2=%d", len(m2))        // 打印 m2 的长度
	m3 := make(map[int]int, 10)         // 使用 make 函数定义一个初始容量为 10 的 map
	t.Logf("len m3=%d", len(m3))        // 打印 m3 的长度
}

// TestAccessNotExistingKey 测试访问不存在的 key
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}     // 定义一个空的 map
	t.Log(m1[1])            // 打印 key 为 1 的值(不存在)
	m1[2] = 0               // 设置 key 为 2 的值为 0
	t.Log(m1[2])            // 打印 key 为 2 的值
	m1[3] = 0               // 设置 key 为 3 的值为 0
	if v, ok := m1[3]; ok { // 检查 key 为 3 是否存在
		t.Logf("Key 3's value is %d", v) // 存在则打印其值
	} else {
		t.Log("key 3 is not existing.") // 不存在则打印提示信息
	}
}

// TestTravelMap 遍历并打印 map 中的所有键值对
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9} // 定义并初始化一个 map
	for k, v := range m1 {              // 遍历 map 中的所有键值对
		t.Log(k, v) // 打印键和值
	}
}
