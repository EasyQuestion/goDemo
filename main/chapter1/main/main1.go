package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

//  错误信息 go run: cannot run non-main package
//  main方法只能放在package main中

// go run main1.go  可以直接运行文件
// go build main1.go  生成main.exe  直接执行main.exe文件
/*   go build会将go文件编译成二进制文件  */
