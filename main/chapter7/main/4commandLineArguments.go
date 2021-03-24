package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argswithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argswithoutProg)
	fmt.Println(arg)
}
// 命令行参数是指定程序运行参数的一个常见方式
// os.Args 提供原始命令行参数访问功能。注意，切片中的第一个参数是该程序的路径，并且 os.Args[1:]保存所有程序的的参数。
// 你可以使用标准的索引位置方式取得单个参数的值
// go build command-line-arguments.go
// ./command-line-arguments a b c d

// 参数化程序的基本方式：命令行参数、命令行标志、环境变量