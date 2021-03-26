package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = false
	fmt.Println("b=", b)
	fmt.Printf("b的类型：%T 占用的字节大小：%d", b, unsafe.Sizeof(b))
}

// bool类型只有true,false，没有对应的0，1..。
// bool类型占1个字节
// boolean类型适于逻辑运算，一般用于程序流程控制（if,for）
