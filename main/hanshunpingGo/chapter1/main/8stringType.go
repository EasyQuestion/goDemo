package main

import "fmt"

func main() {
	var str string = "北京长城"
	fmt.Println(str)

	var str1 = "hello"
	// str1[0] = 'a'  // Cannot assign to str1[0]
	str1 = "world"
	fmt.Println(str1)

	str2 := `
package main

import "fmt"

func main() {
	fmt.Println("hello,world!")
}
`
	//fmt.Println(str2)

	str2 = "hello " + "world"
	str2 += " haha!"
	fmt.Println(str2)

	str2 = "helloasdfdfafdfasfas" + // 换行时，+一定要留在上一行，否则报错（go会默认在行尾加分号，导致编译出错）
		"helloasdfdfafdfasfas"
}

// go中的字符串是由单个字节连接起来的，字节使用的是utf-8编码的unicode文本

// go 中string ,使用注意事项
// 1.go中的字符串中的字节使用utf-8编码标识unicode文本，统一编码方式，不会乱码
// 2.go中字符串一旦赋值，单个字节的内容是不可变的
// 3.字符串的2种表示方式，方一：双引号 方二：反引用(包括特殊字符和换行，可以防止攻击、实现源码输出等)
// 4.字符串拼接 +
// 5.当一行太长，需要换行时，+一定要留在上一行，否则编译报错
