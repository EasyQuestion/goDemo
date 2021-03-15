package main

import "fmt"

func main() {
	fmt.Println("Hello " + "world")
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

//Go 拥有各值类型，包括字符串，整形，浮点型，布尔型等
