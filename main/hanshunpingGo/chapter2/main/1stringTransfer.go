package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "hello"
	num, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("num 的类型是%T num=%v", num, num)
}

/*func main() {
	var str string = "true"
	b, _ := strconv.ParseBool(str)
	fmt.Printf("b 的类型是%T b=%v\n", b, b)

	str2 := "1234590"
	n1, _ := strconv.ParseInt(str2, 10, 32)
	fmt.Printf("n1 的类型是%T n1=%v\n", n1, n1)
	n2 := int32(n1)
	fmt.Printf("n2 的类型是%T n2=%v\n", n2, n2)

	str3 := "23.456"
	n3, _ := strconv.ParseFloat(str3, 64)
	fmt.Printf("n3 的类型是%T n3=%v\n", n3, n3)

	//注意:ParseInt,ParseFloat都是直接转换成64位的结果,如果需要其它精度的类型,需要自己再手动转换一次
}*/

/*func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = false
	//var myChar byte = 'h'
	var str string

	// 10表示十进制展示
	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str 类型是%T,str=%q\n", str, str)

	// 'f'表示浮点数十进制  10表示小数保留10位 64表示是float64(32可能会出现精度丢失情况)
	str = strconv.FormatFloat(num2, 'f', -1, 64)
	fmt.Printf("str 类型是%T,str=%q\n", str, str)

	str = strconv.FormatBool(b)
	fmt.Printf("str 类型是%T,str=%q\n", str, str)

	// 疑问点,byte类型怎么转换成string?
	//str = strconv.FormatUint(uint64(myChar), 10)
	//fmt.Printf("str 类型是%T,str=%q(%c)\n", str, str, str)
	//str = strconv.Itoa(int(myChar))
	//fmt.Printf("str 类型是%T,str=%q\n", str, str)

	str = strconv.Itoa(num1)
	fmt.Printf("str 类型是%T,str=%q\n", str, str)

}*/

/*func main() {

	var num1 int = 99
	var num2 float32 = 23.456
	var b bool = false
	var myChar byte = 'h'
	var str string

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str 类型是%T,str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str 类型是%T,str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str 类型是%T,str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str 类型是%T,str=%q\n", str, str)
}*/

// 基本数据类型与string之间的转换,2种方式
// 方一：fmt.Sprintf("%参数",表达式)  很常用
// 方二：strconv包中的函数

// string转换成基本数据类型
// strconv包中的函数

// 注意事项:string转换成基本数据类型时,要确保string类型能够转成有效的数据
// "hello"如果转换成int类型,并不会报错,只是转换的结果是0(转换成了默认值)
