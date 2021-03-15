package main

import "fmt"

/*func plus(a int, b int) int {
	return a + b
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)
}*/
// 函数 是 Go 的中心(特性：1.多值返回  2.允许可变长参数 3.闭包结构-匿名函数 4.递归)

// 这里有许多 Go 函数的其他特性。其中一个就是多值返回
// Go 内建多返回值 支持。这个特性在 Go 语言中经常被用到，例如用来同时返回一个函数的结果和错误信息。
// (int, int) 在这个函数中标志着这个函数返回 2 个 int。
// 通过多赋值 操作来使用这两个不同的返回值。
// 仅仅想返回值的一部分的话，你可以使用空白定义符 _。
/*func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a, b)
	_, c := vals()
	fmt.Println(c)
}*/

// 允许可变长参数是 Go 函数的另一个好的特性
// 可以用任意数量的参数调用。例如，fmt.Println 是一个常见的变参函数
// 如果你的 slice 已经有了多个值，想把它们作为变参使用，你要这样调用 func(slice...)。
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
