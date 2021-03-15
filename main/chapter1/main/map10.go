package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	// 内建的 delete 可以从一个 map 中移除键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	// 当从一个 map 中取值时，可选的第二返回值指示这个键是在这个 map 中。
	// 这可以用来消除键不存在和键有零值，像 0 或者 "" 而产生的歧义。
	_, prs2 := m["k2"]
	fmt.Println("prs:", prs2)

	// 可以通过这个语法在同一行申明和初始化一个新的map。
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
// 要创建一个空 map，需要使用内建的 make
