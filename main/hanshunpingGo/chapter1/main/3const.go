package main

import "fmt"

func main() {
	num := 9
	const c = 9 / 3
	//const c = num / 3 // num是变量，所以不能编译器不能通过，编译时必须要确定的值
	//const c = getVal() 这种编译器也不能通过
	fmt.Println(num, c)

	// const简洁的写法
	//const (
	//	a = 1
	//	b = 2
	//)

	//const 专业写法
	//const (
	//	x = iota
	//	y
	//	z
	//)
	//fmt.Println(x, y, z) // 0,1,2

	const (
		x    = iota
		y    = iota
		m, n = iota, iota
	)
	fmt.Println(x, y, m, n) // 0,1,2,2
}

// 常量 const
// 在定义时必须初始化(没有默认值)
// 常量不能修改
// const 只能修饰 bool,string,数值类型(int,float系列)
// 语法 const xxx变量名 [type] = xxx值  // type可以省略

// const 使用注意事项
// 1.常量没必要全部大写
// 2.常量仍通过首字母大写控制访问范围
