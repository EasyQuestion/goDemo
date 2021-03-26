package main

import "fmt"

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'
	// 直接输出，会输出对应的ASCII码值
	fmt.Println("c1=", c1, "c2=", c2)
	// 格式化输出即可
	fmt.Printf("c1=%c,c2=%c \n", c1, c2)

	//var c3 byte = '北' // constant 21271 overflows byte
	var c3 int = '北' // utf-8编码  go中只使用这个编码  进行unicode编码（其中的一种具体实现就是utf-8）
	fmt.Printf("c3=%c 对应的码值：%d \n", c3, c3)

	var c4 int = 22269 // -> 国
	fmt.Printf("c4=%c", c4)

	var n1 = 10 + 'a'// -> 10 + 97 = 107
	fmt.Println("n1=", n1)
}

/*func main() {
	var n1 = 100
	fmt.Printf("n1 的类型是 %T \n", n1)

	var n2 int64 = 100
	fmt.Printf("n2 的类型是 %T ,n2占用字节大小：%d", n2, unsafe.Sizeof(n2))
}*/

// 整数使用，注意事项
// 1.各整数类型分 有符号 和 无符号
// 2.整数默认会匹配成 int 类型
// 3.看某个变量的数据类型和字节大小
// 4.整型变量使用原则：保小不保大（即在保证程序正常运行下，尽量使用占用空间小的数据类型）
// 5.bit 计算机中最小存储单位，byte 计算机中基本存储单元 1 byte = 8 bit

// go 中，没有专门的字符类型，如果要存储单个字符（字母：ASCII码），一般用byte来保存
// 传统字符串是由字符组成的，而go中字符串是由字节组成的

// 字符 byte使用，注意事项
// 1.字符常量是用单引号括起来的单个字符
// 2.go中允许使用转义字符表示特殊字符常量
// 3.GO中字符使用utf-8编码，英文字母1个字节，汉字3个字节
// 4.字符的本质是一个整数，直接输出时，是该字符对应的utf-8编码值
// 5.可以直接给某个变量赋一个数字，然后按格式化%c输出该数字对应的unicode字符
// 6.字符类型可以进行运算，相当于一个数字，因为它都对应有unicode码

// 字符类型的本质讨论：字符 -> utf-8码值 -> 二进制 -> 存储
