package testing

import (
	"testing" // 导入测试包

	. "github.com/smartystreets/goconvey/convey" // 导入goconvey包
)

func TestSpec(t *testing.T) {

	// 仅传递 t 参数到顶级的 Convey 调用中
	Convey("Given 2 even numbers", t, func() { // 给定 2 个偶数
		a := 3 // 定义变量 a 为 3
		b := 4 // 定义变量 b 为 4

		Convey("When add the two numbers", func() { // 当添加这两个数字时
			c := a + b // 定义变量 c 为 a 和 b 之和

			Convey("Then the result is still even", func() { // 那么结果仍然是偶数
				So(c%2, ShouldEqual, 0) // 断言 c 对 2 取余等于 0
			})
		})
	})
}
