package main

import "fmt"

func main() {
	fmt.Println("10/4=", 10/4)
	fmt.Println("10.0/4=", 10.0/4)

	fmt.Println("10%3=", 10%3)
	fmt.Println("-10%3=", -10%3)
	fmt.Println("10%-3=", 10%-3)
	fmt.Println("-10%-3=", -10%-3)

	i := 10
	i++
	fmt.Println("i=", i)
	i--
	fmt.Println("i=", i)
}

// 运算符:
// 1.算术运算符(对数值型变量运算)+ - * / % a++ a-- a+b
// 2.赋值运算符
// 3.比较运算符
// 4.逻辑运算符
// 5.位运算符
// 6.其它运算符
