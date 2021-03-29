package main

import "fmt"

func main() {
	var a int
	var b float32
	var c float64
	var isMarried bool
	var str string
	fmt.Println("a=", a, " b=", b, " c=", c, " isMarried=", isMarried, " str=", str)
	fmt.Printf("a=%d,b=%v,c=%v,isMarried=%v,str=%v", a, b, c, isMarried, str)
	// %v 表示按原始值输出
}

//func main() {
//	var i int = 1
//	fmt.Println("i=", i)
//
//	//var j int8 = 128//-129
//	//fmt.Println("j=", j)
//
//	var j uint8 = 256 //-1
//	fmt.Println("j=", j)
//}

// 变量的数据类型：基本数据类型、派生/复杂数据类型
// 基本数据类型：数值型（整数类型、浮点类型、byte）、布尔型、字符串
// 派生/复杂数据类型（8个）：数组、切片、map、指针、结构体、管道、函数、接口

// 有符号的整数 int8:-128~127 (-2^7 ~ 2^7 -1)
// 无符号的整数 uint8:0~255 (0 ~ 2^8-1)

// int  默认是有符号的，到底表示几位和系统有关（32位系统就表示32位即4个字节，64位系统就表示64位即8个字节）
// uint 默认是无符号的，到底表示几位和系统有关
// rune 有符号 与int32一样，表示一个unicode码 （一般汉字都用这个类型）
// byte 无符号 与uint8等价  0~255 （当要存储字符时，选用byte）

// 基本数据类型的默认值（也叫零值）
// 整型 0  浮点型 0  字符串 ""  布尔型 false
