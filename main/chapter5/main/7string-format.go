package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	//p := point{1, 2}
	//fmt.Printf("%v\n", p) // {1 2}
	//fmt.Printf("%+v\n", p) // {x:1 y:2}
	//fmt.Printf("%#v\n", p)// main.point{x:1, y:2}
	//fmt.Printf("%T\n", p) // main.point  值的类型
	//fmt.Printf("%t\n", true) // true 格式化布尔值
	//fmt.Printf("%d\n",123) // 十进制格式
	//fmt.Printf("%b\n",14)// 二进制
	//fmt.Printf("%c\n", 33)//!  给定整数的对应字符
	//fmt.Printf("%x\n", 456) // 1c8 十六进制编码
	//fmt.Printf("%f\n", 78.9)// 78.900000 对于浮点型进行最基本的十进制格式化

	// %e 和 %E 将浮点型格式化为（稍微有一点不同的）科学技科学记数法表示形式。
	//fmt.Printf("%e\n", 123400000.0)
	//fmt.Printf("%E\n", 123400000.0)

	//fmt.Printf("%s\n", "\"string\"") // "string" 字符串输出
	//fmt.Printf("%q\n", "\"string\"")// "\"string\""  源代码中那样带有双引号的输出
	//fmt.Printf("%x\n", "hex this") //6865782074686973   %x 输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示
	//fmt.Printf("%p\n", &p)// 0xc042050070   输出一个指针的值

	//fmt.Printf("|%6d|%6d|\n", 12, 345)       // 在 % 后面使用数字来控制输出宽度。默认结果使用右对齐并且通过空格来填充空白部分。
	//fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) // 可以通过 宽度.精度 的语法来指定输出的精度
	//fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // 要左对齐，使用 - 标志

	//fmt.Printf("|%6s|%6s|\n", "foo", "b") // 控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。这是基本的右对齐宽度表示。
	//fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // 要左对齐，和数字一样，使用 - 标志。

	//s := fmt.Sprintf("a %s", "string") // 返回一个字符串
	//fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error") // 格式化并输出到 io.Writers而不是 os.Stdout
}
