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

//func sumFunc(args ...int) (sum int) {
/*for i := range args { //???疑问：这个不会取最后一个值，为啥
	fmt.Println("i=",i)
	sum += i
}
return*/

//	for i := 0; i < len(args); i++ {
//		sum += args[i]
//	}
//	return
//}
//
//func main() {
//	fmt.Println("sum=", sumFunc(15))
//}

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
//----------------------------------------------------
//  执行顺序：全局变量定义 -> init函数 -> main函数
/*var age = test()

func test() int {
	fmt.Println("test()...")
	return 90
}

// 每个源文件都包含一个Init函数，在main函数执行前被go调用
func init() {
	fmt.Println("init()...")
}

// 先执行的是引入包的init函数，然后是本文件的全局变量定义，本文件的init函数，本文件的main函数
func main() {
	fmt.Println("main()...")
	fmt.Println(utils.Age, utils.Name)
}*/

//----------------------------------------------------------
/*var (
	// 全局匿名函数
	Fun1 = func(n1, n2 int) int {
		return n1 * n2
	}
)

func main() {
	// 方一：匿名函数直接调用
	result := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("result=", result)

	// 方二：将匿名函数交给变量，通过变量来实现多次调用
	a := func(n1, n2 int) int {
		return n1 - n2
	}

	result2 := a(20, 10)
	fmt.Println("result2=", result2)

	result3 := Fun1(4, 9)
	fmt.Println("result3=", result3)
}*/
//----------------------------------------------
// 方法体内的代码形成一个闭包
// 闭包相当于一个类，n是字段，匿名函数是操作
/*func AddUpper() func(int) int {
	// 返回一个匿名函数，匿名函数引用到了函数外的n，因此这个匿名函数和n就形成一个整体，构成闭包
	var n = 10
	str := "hello"
	return func(x int) int {
		n += x
		str += fmt.Sprintf("%d", x)
		fmt.Println(str)
		return n
	}
}

// 要求：判断文件后缀名是否是以.jpg结尾，如果是就原样返回，如果不是就加上
// 闭包的好处：普通方法也可以实现，但闭包实现可以不用多次输入后缀
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			name += ".jpg"
		}
		return name
	}
}

func main() {
	f := AddUpper()
	// 当反复调用f时，由于n只初始化一次，因此每调用一次就累加一次
	fmt.Println(f(1))
	fmt.Println(f(2))
	f2 := AddUpper()
	fmt.Println(f2(2))
	fmt.Println(f2(3))

	f3 := makeSuffix(".jpg")
	fmt.Println("处理后的文件名：", f3("window"))// window.jpg
	fmt.Println("处理后的文件名：", f3("bird.jpg"))// bird.jpg

}*/
//----------------------------------------------------------
// 当执行到defer语句时，会将语句先压入栈中（一个独立的栈，可以认为是defer栈）
// 等函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行语句
/*func sumDefer(n1 int, n2 int) int {
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

func main() {
	res := sumDefer(10, 20)
	fmt.Println("res=", res)
}
*/
//----------------------------------------------------------------
// 全局变量：可以全局定义变量，也可以全局初始化一个变量，但是不能有赋值语句
var name string
// 全局变量使用陷阱，不能这样使用
// := 相当于2个语句 1)var Name string  2)Name = "jack"  1)可以编译通过，但2)不能编译通过
//name = "tome"
//Name := "jack"
func main() {
	fmt.Println("name=", name)
	//fmt.Println("Name=",Name)
}

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

// go vendor
// 基本思路是将引用的外部包的源代码放在当前工程的 vendor 目录下面
// 编译 go 代码会优先从 vendor 目录先寻找依赖包
// 有了 vendor 目录后，打包当前的工程代码到其他机器的$GOPATH/src 下都可以通过编译

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

// init初始化函数
//  每个源文件都包含一个Init函数，在main函数执行前被go调用
//  通常可以在init函数中完成初始化的工作
//  执行顺序：全局变量定义 -> init函数 -> main函数
//  当引入其它包文件变量时，先是引入包内的全局变量定义->引入包的init函数-> 本包内的init函数 -> 本包的main函数

// 匿名函数
// 使用方式：1.定义时直接使用 2.将匿名函数交给一个变量（函数变量），通过变量来调用
// 如果将匿名函数交给一个全局变量，那么这个匿名函数就成为一个全局匿名函数，可以在整个程序有效

// 闭包：一个函数与其相关的引用环境组合的一个整体（实体）
// 返回一个匿名函数，匿名函数引用到了函数外的n，因此这个匿名函数和n就形成一个整体，构成闭包
// 闭包相当于一个类，n是字段，匿名函数是操作
// 闭包的关键：返回的函数和它引用到的变量，两者共同构成了闭包

// defer关键字：延时机制
// 创建资源（数据库连接、文件句柄、锁等），为了在函数执行完毕后，及时释放资源，GO的设计者提供defer(延时机制)
// 当执行到defer语句时，会将语句先压入栈中（一个独立的栈，可以认为是defer栈）此时可以继续使用相关的资源
// 等函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行defer语句
// 入栈时，会将相关的值拷贝同时入栈，所以要注意
// 这样，程序员就不用再烦恼什么时候关闭资源了

// 函数的参数传递方式（值传递、引用传递）重点
// 一般来说引用传递效率高些，因为传的是地址，数据量小，而值拷贝取决于数据大小（比如数组，结构体），要拷贝的数据越大，效率越低
// 如果要修改原变量的值，需要用指针&

// 变量作用域
// 1.在函数内定义或声明的变量，作用域在函数内
// 2.函数外定义或声明的变量，作用域在整个包有效，如果其首字母大写，则作用域在整个程序有效
// 3.如果变量是在代码块，如if/for/switch中，那么变量作用域就在该代码块
// 全局变量：可以全局定义变量，也可以全局初始化一个变量，但是不能有赋值语句

// 函数 vs 方法
