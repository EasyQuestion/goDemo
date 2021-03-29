package main

import "fmt"

func main() {

	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	var str = i
	fmt.Printf("i=%v n1=%v n2=%v n3=%v str=%v\n", i, n1, n2, n3, str)
	fmt.Printf("i 的类型为%T \n", i)

	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=", num2)
}

// 数据类型转换：go中只能显示转换， 不能自动转换
// 语法格式：T(v)

// 说明
// 1.go中数据类型的转换，包括 低精度->高精度 高精度->低精度 数值范围小->大 大->小
// 2.被转换的是变量存储的值，变量本身的数据类型并没有变量
// 3.int64 -> int8 编译时不会报错，只是转换的结果按溢出处理，和希望的结果不一样，因此在转换时，需要考虑范围
