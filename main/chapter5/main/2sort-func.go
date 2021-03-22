package main

import (
	"fmt"
	"sort"
)

type ByLength []string

// 我们在类型中实现了 sort.Interface 的 Len，Less和 Swap 方法，这样我们就可以使用 sort 包的通用Sort 方法了
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
// 按字符串长度增加的顺序来排序
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	// 通过将原始的 fruits 切片转型成 ByLength 来实现我们的自定排序了
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
// 类似的，参照这个创建一个自定义类型的方法，实现这个类型的这三个接口方法，
// 然后在一个这个自定义类型的集合上调用 sort.Sort 方法，我们就可以使用任意的函数来排序 Go 切片了。