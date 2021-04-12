package main

import (
	"errors"
	"fmt"
)

/*func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2 // 此时抛出的异常叫做panic
	fmt.Println("res=", res)
}

func main() {
	test()
	fmt.Println("end...")
}*/
//------------------------------------
func readConf(name string) (err error) {
	if name == "init.conf" {
		return nil
	} else {
		return errors.New("读取文件错误...")
	}
}

func test() {
	//err := readConf("init.conf")
	err := readConf("start.conf")
	if err != nil {
		// 打印这个错误并停止程序
		panic(err)
	}
	fmt.Println("other process..")
}

func main() {
	// 测试自定义错误
	test()
	fmt.Println("end..")
}

// 异常处理机制：defer,panic,recover
// go中可以抛出panic异常，在defer中通过recover捕获这个异常，之后进行正常的流程
// 自定义错误 errors.New("错误信息")，panic内置函数(接收error类型的变量，打印并退出程序)
