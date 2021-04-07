/*package main // 打包
import (
	"fmt"                                      // 导jdk包
	//f "fmt"                                    // 对导入的包起别名
	"goDemo/main/hanshunpingGo/chapter3/utils" // 导自定义的包（即导入函数所在的文件夹路径，默认从%GOPATH%/src/下寻找）
)

func main() {

	fmt.Println("hello")       // 通过“包名.函数名”来调用包内的函数
	fmt.Println(abc.Cal(2, 3)) // 通过“包名.函数名”来调用包内的函数

}*/
//---------------------------------------------------------------------
package main

import "fmt"

// 函数的调用及内存分配（值类型在栈中，值传递，但是内存地址不同）
/*func test(n1 int) {
	n1 += 1
	fmt.Println("test() n1:", &n1, "n1=", n1)
}

func main() {
	n1 := 10
	test(n1) // 函数调度，此时 main()中的n1仍为10
	fmt.Println("main() n1:", &n1, "n1=", n1)
}*/
//-------------------------------------------------
/*func test(n1 int, n2 int) (int, int) { // 多个返回值
	sum := n1 + n2
	suff := n1 - n2
	return sum, suff
}*/

/*func test(n1 int, n2 int) (sum int, sub int) { // 支持对返回值命名，多个返回值不用再考虑顺序问题
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	// 同时计算两个数的和，差
	n1 := 12
	n2 := 8
	sum, sub := test(n1, n2)
	fmt.Println("sum=", sum, "sub=", sub)
}*/

func sumFunc(args ...int) (sum int) {
	/*for i := range args { //???疑问：这个不会取最后一个值，为啥
		fmt.Println("i=",i)
		sum += i
	}
	return*/

	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

func main() {
	fmt.Println("sum=", sumFunc(15))
}

//------------------------------------------------
/*func test(n int) { // 递归：自己调自己
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=", n)
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("n=", n)
	}
}

func main() {
	test(4)
	test2(4)
}*/
//---------------------------------------------------------------
/*func getSum(n1 int, n2 int) int {
	return n1 + n2
}*/

/*func myFunc(f func(int, int) int, num1 int, num2 int) int {
	return f(num1, num2)
}*/

/*type mySum func(int, int) int

func myFunc2(f mySum, num1 int, num2 int) int {
	return f(num1, num2)
}*/

/*
func main() {
	// 函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量，通过该变量调用函数
	a := getSum
	fmt.Printf("a的变量类型是%T\ngetSum的变量类型是%T\n", a, getSum)

	res := a(4, 2)
	fmt.Println("res=", res)

	res2 := myFunc(getSum, 10, 30)
	fmt.Println("res2=", res2)
}*/
//---------------------------------------------------------
/*func main() {
	type myInt int // 从语法上看，myInt,int是2个不同的数据类型
	var num myInt
	num = 40
	fmt.Printf("num的类型是%T num=%v\n", num, num)
	var num2 int
	//num2 = num // 编译不通过，因为类型不同
	num2 = int(num)
	fmt.Println("num2=", num2)

	res := myFunc2(getSum, 11, 22)
	fmt.Println("res=", res)
}*/

/* 函数的基本语法
func 函数名（形参列表）（返回值列表）{
	执行语句
	return 返回值列表
}
// 返回值列表可以没有
// 支持对函数返回值命名
// 可以用 _ 来忽略返回值
// 支持可变参数 语法
func sum(args... int)sum int {
}

func sum(n1 int,args... int)sum int {
}
args类型是slice切片
*/

// 包：将变量和函数分散到不同文件中去，也可以解决变量和函数重名的问题，还可以控制变量和函数的访问范围
// （本质就是文件夹）
// 函数名大写，表示其它包能访问，也叫“该函数可导出”
// 注：包名和文件夹名可以不同，导包的时候用文件夹路径，引函数的时候用”包名.函数名“，但是一般情况下保持一致
// 支持对导入的包起别名
// 同一包下函数名不能重复
// 可执行文件所在的包必须叫main,只能有一个，库文件包可以有多个
// （正规编译）在%GOPATH%路径下，go build -o hello.exe goDemo\main\hanshunpingGo\chapter3\main 编译成可执行文件
// 在生成可执行文件时，也会生成工具包的库文件例如utils.a，可以把库文件交给别人，直接用
// ？？？疑问：生成可执行文件时，没有生成对应的库文件，而且可执行文件只能在终端执行，不能双击运行

//
// 函数的调用及内存分配（值类型在栈中，值传递，但是内存地址不同，如果要修改原变量的值，需要用指针&）
// 递归调用(注意：递归必须有出口，否则会死循环)
// go中不支持重载（因为在go中支持可变参数，空接口）
// 函数内的局部变量，作用域范围只在函数内
// 函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量，通过该变量调用函数
// 函数是一种数据类型，函数可以作为形参，并且调用
// 为了简化数据类型定义，go支持自定义数据类型
/*基本语法
type 自定义数据类型名 数据类型
*/

// 函数 vs 方法
