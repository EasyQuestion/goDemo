package main

import "fmt"

func main() {
	/*nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}*/

	/*kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}*/

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
// range 迭代各种各样的数据结构
// 使用 range 来统计一个 slice 的元素和。数组也可以采用这种方法。
// range 在数组和 slice 中都同样提供每个项的索引和值.
// 不需要索引，所以我们使用 空值定义符_ 来忽略它
// range 在 map 中迭代键值对。
// range 在字符串中迭代 unicode 编码。