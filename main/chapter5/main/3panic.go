package main

import "os"

func main() {
	//panic("a problem")

	_, err := os.Create("/tmp/file")
	// 创建一个新文件时返回异常错误时的panic 用法
	if err != nil {
		panic(err)
	}
}

// panic 的一个基本用法就是在一个函数返回了错误值但是我们并不知道（或者不想）处理时终止运行。
