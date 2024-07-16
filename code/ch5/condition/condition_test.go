package condition_test

import "testing"

// TestSwitchMultiCase 函数: 测试多 case switch 语句
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even") // 如果 i 为 0 或 2: 打印 "Even"
		case 1, 3:
			t.Log("Odd") // 如果 i 为 1 或 3: 打印 "Odd"
		default:
			t.Log("it is not 0-3") // 如果是其他数值: 打印 "it is not 0-3"
		}
	}
}

// TestSwitchCaseCondition 函数: 测试带条件判断的 switch 语句
func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even") // 如果 i 是偶数: 打印 "Even"
		case i%2 == 1:
			t.Log("Odd") // 如果 i 是奇数: 打印 "Odd"
		default:
			t.Log("unknow") // 如果是其他情况: 打印 "unknow"
		}
	}
}
