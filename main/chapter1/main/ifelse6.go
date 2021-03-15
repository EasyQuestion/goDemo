package main

import "fmt"

func main() {
	/*if 7%2 == 0 {
		fmt.Println("7 is even")
	}else{
		fmt.Println("7 is odd")
	}*/

	/*if 8%4 == 0{
		fmt.Println("8 is divisible by 4")
	}*/

	// 在条件语句之前可以有一个语句；任何在这里声明的变量都可以在所有的条件分支中使用
	if num := 9; num < 0 {
		fmt.Println("is negative")
	} else if num < 10 {
		fmt.Println("has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// 在 Go 中，你可以不适用圆括号，但是花括号是需要的。
// Go 里没有三目运算符，所以即使你只需要基本的条件判断，你仍需要使用完整的 if 语句。