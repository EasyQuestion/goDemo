package main

import "fmt"

func main() {
	var a interface{}
	var num = 8.9
	a = num
	fmt.Println(a)

	// 带检测的类型断言
	if b, ok := a.(float32); ok {
		fmt.Println("convert success,b=", b)
	} else {
		fmt.Println("convert failure")
	}

	fmt.Println("继续执行...")
}

/*type Point struct {
	x, y int
}

func main() {
	p := Point{1, 2}
	var a interface{} = p
	fmt.Println(a)
	var b Point
	// b = a
	// 类型断言，如果a是指向Point类型的变量，就转成Point类型，并赋值给b，否则报错
	b = a.(Point)
	fmt.Println(b)
}*/
