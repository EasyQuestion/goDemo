package main

import "fmt"

func main() {
	var price float32 = 89.12
	fmt.Println("price=", price)

	var num1 float32 = -0.00089
	var num2 float64 = -7809695.09
	fmt.Println("num1=", num1, "num2=", num2)

	var num3 float32 = -123.000901
	var num4 float64 = -123.000901
	fmt.Println("num3=", num3, "num4=", num4)

	var num5 = 1.1
	fmt.Printf("num5 的类型为 %T \n", num5)

	num6 := 5.12
	num7 := .123
	fmt.Println("num6=", num6, "num7=", num7)

	num8 := 5.1234e2 // 相当于 5.1234 * 10^2
	num9 := 5.1234E2
	num10 := 5.1234e-2 // 相当于 5.1234 / 10^2
	fmt.Println("num8=", num8, "num9=", num9, "num10=", num10)
}

// 浮点点 = 符号位 + 指数位 + 尾数位  (浮点数都是有符号的)
// 尾数部分有可能丢失，造成精度损失

// 浮点类型使用，注意事项
// 1.浮点类型有固定的范围和字段长度，不受具体OS（操作系统）的影响
// 2.浮点类型默认为float64类型
// 3.浮点类型表示方式 方一：十进制  方二：科学计数法
// 4.通常情况下，应该使用float64，因为它比float32更精确
