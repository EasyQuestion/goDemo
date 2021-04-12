package main

import (
	"fmt"
	"strings"
)

func main() {

	//fmt.Println(len("中")) // 3

	/*str := "hello北"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}*/

	/*str, err := strconv.Atoi("hello") // "12"
	if err != nil {
		fmt.Println("转换错误：", err)
	} else {
		fmt.Printf("str的类型是%T,str=%v", str, str)
	}*/

	//str := fmt.Sprintf("%c",'a')
	//fmt.Printf("str的类型是%T,str=%v", str, str)
	/*str := strconv.Itoa(12)
	fmt.Printf("str的类型是%T,str=%v", str, str)*/

	/*bytes := []byte("hello go")
	fmt.Printf("bytes的类型是 %T",bytes)*/

	/*str := string([]byte{97, 98, 99})
	fmt.Printf("str的类型是%T,str=%v", str, str)*/

	/*intTest := 8
	fmt.Printf("十进制：%d 二进制：%b 八进制：%o 十六进制：%x", intTest, intTest, intTest, intTest)
	str := strconv.FormatInt(8, 2)
	fmt.Printf("str的类型是%T,str=%v", str, str)*/

	//fmt.Println(strings.Contains("hello go go", "o g"))

	//fmt.Println(strings.Count("hellofooeoocoo", "cc"))

	/*fmt.Println(strings.EqualFold("hello", "HELLO"))//不区分大小写
	fmt.Println("hello" == "HELLO")//区分大小写*/

	//fmt.Println(strings.Index("hello", "e"))
	//fmt.Println(strings.LastIndex("helloll", "ll"))

	/*str := "golang"
	str = strings.Replace(str, "go", "to", 2) // -1 代表全部 1代表替换1个
	fmt.Println(str)*/

	//str := strings.Split("hello,work,ok", ",")
	//fmt.Println(str)

	/*str := "hello,Abc"
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))*/

	/*str := " go go hello   "
	fmt.Printf("%q", strings.TrimSpace(str))*/

	/*str := "! hello!go!! !"
	fmt.Println(strings.Trim(str, "! "))//会把左右两边的空格和叹号都去掉
	fmt.Println(strings.TrimLeft(str, "! "))
	fmt.Println(strings.TrimRight(str, "!"))*/

	str := "go go hello"
	fmt.Println(strings.HasPrefix(str, "go"))
	fmt.Println(strings.HasSuffix(str, "ello"))
}

// string中常用的函数
// 1.统计string长度 len(str)  查看文档手册：内建函数builtin
// 2.string遍历时，如果含有中文，需要先转换成[]rune(str),再进行遍历，或者直接用range进行遍历
// 3.string转int strconv.Atoi("12")
// 4.int转string strconv.Itoa(12)
// 5.string转[]byte切片
// 6.[]byte切片转string
// 7.十进制转二、八、十六进制 strconv.FormatInt(8, 2)
// 8.查找string中是否包含某个子串 strings.Contains("hello go go", "o g")
// 9.统计在string中，某个子串出现的次数 strings.Count("hellofooeoocoo", "oo")
// 10.忽略大小写的string比较 方一：==（区分大小写） 方二：strings.EqualFold("hello", "HELLO") //不区分大小写
// 11.返回某子串在string中第一次出现的下标 strings.Index("hello", "e")
// 12.返回某子串在string中最后一次出现的下标 strings.LastIndex("helloll", "ll")
// 13.将string中指定的子串替换 strings.Replace(str, "go", "to", 1)// -1 代表全部 1代表替换1个
// 14.将string拆分成数组 strings.Split("hello,work,ok", ",")
// 15.string大小写转换 strings.ToUpper(str) strings.ToLower(str)
// 16.将string两边的空格去掉 strings.TrimSpace(str)
// 17.将string两边指定的字符串去掉 strings.Trim(str, "!")
// 18.将string左边指定的字符串去掉 strings.TrimLeft(str, "!")
// 19.将string右边指定的字符串去掉 strings.TrimRight(str, "!")
// 20.判断string是否以某子串开头 strings.HasPrefix(str, "go")
// 21.判断string是否以某子串结尾 strings.HasSuffix(str, "ello")
