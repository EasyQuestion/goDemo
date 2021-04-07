package main

import "fmt"

func main() {
	var num int
	fmt.Println("num=", num, "num的地址:", &num)

	ptr := &num
	*ptr = 10
	fmt.Println("num=", num, "ptr=", ptr)
}

/*func main() {
	var i int = 10
	fmt.Println("i的内存地址:", &i)

	var ptr *int = &i
	fmt.Println("ptr=", ptr, "\n&ptr=", &ptr)
	fmt.Println("ptr指向的值:", *ptr, "->", *&i)
}*/

// 指针
// 1.基本数据类型,变量存的就是值,也叫值类型
// 2.获取变量的地址,用&,比如:var num int,获取num的地址,&num,了解基本数据类型在内存中的布局
// 3.指针类型,指针变量存的是一个地址,这个地址指向的空间存的才是真正的值
// 4.获取指针类型所指向的值,使用*,比如var ptr *int,使用*ptr获取地址真正存的值

// 指针的使用细节
// 1.值类型,都有对应的指针类型,语法为*数据类型,比如*int,*float...
// 2.值类型包括:基本数据类型 int系列,float系列,bool,string,数组和结构体struct

// 值类型 vs 引用类型
// 值类型:基本数据类型 int系列,float系列,bool,string,数组和结构体
// 引用类型:指针,slice切片,map,管道chan,interface等都是引用类型

// 值类型,变量直接存储值,内存通常在栈中分配(go中有编译逃逸分析,也有可能在堆区)
// 引用类型,变量存储的是一个地址,地址对应的空间才真正存储值,内存通常在堆上分配,(也有可能在栈区)
// 当没有任何变量引用这个地址时,该地址对应的数据空间就成为一个垃圾,由GC回收
