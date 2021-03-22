package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))      // peach 查找匹配字符串
	fmt.Println(r.FindStringIndex("peach punch")) //[0 5]  查找第一次匹配的字符串的，但是返回的匹配开始和结束位置索引

	fmt.Println(r.FindStringSubmatch("peach punch"))      // Submatch 返回完全匹配和局部匹配的字符串。例如，这里会返回 p([a-z]+)ch 和 `([a-z]+) 的信息
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]

	fmt.Println(r.FindAllString("peach punch pinch", -1)) //[peach punch pinch] 带 All 的这个函数返回所有的匹配项，而不仅仅是首次匹配项。
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2)) //[peach punch] 限制匹配次数 2次

	fmt.Println(r.Match([]byte("peach"))) //true

	// 创建正则表示式常量时，可以使用 Compile 的变体MustCompile 。因为 Compile 返回两个值，不能用于常量。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit> 替换部分字符串为其他值

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)// Func 变量允许传递匹配内容到一个给定的函数中
	fmt.Println(string(out))
}
