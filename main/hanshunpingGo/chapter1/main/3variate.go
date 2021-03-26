package main

import "fmt"

//一次性声明多个全局变量
var (
	n4    = 200
	name2 = "jack"
	n5    = 8
)

func main() {
	// 一次性声明多个变量
	// var n1, n2, n3 int
	// fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)

	// 多个不同类型的变量
	// var n1, name, n3 = 100, "tom", 88
	// fmt.Println("n1=", n1, "name=", name, "n3=", n3)

	// 也可以使用:=方式初始化多个变量
	n1, name, n3 := 100, "tom", 88
	fmt.Println("n1=", n1, "name=", name, "n3=", n3)

	fmt.Println("n4=", n4, "name2=", name2, "n5=", n5)

	// + 只能做同一类型的拼接或相加
	//var i, j = "tome", 23
	//fmt.Println("i+j=", (i + j))
}

// func main() {
// 	// golang变量的3种使用方式：
// 	// 1.指定变量类型，声明后若不赋值，使用默认值
// 	var i int
// 	fmt.Println("i=",i)

// 	// 2.省略类型，根据值自行判断变量类型（类型推导）
// 	var num = 10.11
// 	fmt.Println("num=",num)

// 	// 3.省略var关键字和类型，:= 方式（注意：左侧的变量不能是已声明过的，否则会导致编译错误）
// 	name := "tom"
// 	fmt.Println("name=",name)
// }

// func main() {
// 	var i int = 10
// 	fmt.Println("i=",i)

// 	var j int
// 	j = 10
// 	fmt.Println("j=",j)
// }

// 变量使用的注意事项：
// 1.变量表示内存中的一块存储区域，
// 2.该区域有自己的名称（变量名）和类型（数据类型）
// 3.golang变量的3种使用方式
// 4.多变量声明,多全局变量声明
// 5.该区域的数据值可以在同一类型范围内不断变量，不能改变数据类型
// 6.变量在同一作用域（在一个函数或代码块）内不能重名
// 7.变量3要素：变量名、值、数据类型
// 8.变量如果没有赋初值，编译器会使用默认值（int 0,string 空串）

// 名词概念：变量的声明、初始化、赋值
// +的使用：可以作数值型的加法运算，也可以作字符串拼接
