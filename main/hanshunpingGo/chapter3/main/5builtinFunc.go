package main

import "fmt"

func main() {
	num1 := 100
	fmt.Printf("num1的类型:%T num1的值:%v num1的地址:%v\n", num1, num1, &num1)

	num2 := new(int)
	*num2 = 100
	fmt.Printf("num2的类型:%T num2的值:%v num2的地址:%v 指向的值是%v\n", num2, num2, &num2, *num2)
}

// 内置函数 (内建函数)：为了方便使用，提供了一些函数，不用引包可以直接使用，这些函数叫做内置函数
// len()
// new() 用来分配内存。主要是用来分配值类型，返回的是指针
// make()用来分配内存。主要是用来分配引用类型，比如slice,map,chan
