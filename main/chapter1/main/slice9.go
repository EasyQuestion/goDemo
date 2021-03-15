package main

import "fmt"

func main() {
	// slice 的类型仅由它所包含的元素决定（不像数组中还需要元素的个数）。
	// 要创建一个长度非零的空slice，需要使用内建的方法 make。
	s := make([]string, 3)
	//fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	//fmt.Println("set:",s)
	//fmt.Println("get:",s[2])
	//fmt.Println("len:",len(s))

	// slice 支持比数组更多的操作。
	// 其中一个是内建的 append，它返回一个包含了一个或者多个新值的 slice。
	// 注意我们接受返回由 append返回的新的 slice 值。
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	//fmt.Println("len:",len(s))

	c := make([]string, len(s))
	copy(c, s)
	/*fmt.Println("cpy:", c)
	s[0] = "12"
	fmt.Println("cpy1:",c)
	fmt.Println("s:",s)*/

	// Slice 支持通过 slice[low:high] 语法进行“切片”操作,注意左闭右开[low:high)
	/*l := s[2:5]
	fmt.Println("sl1:",l)
	l = s[2:]
	fmt.Println("sl2:",l)
	l = s[:5]
	fmt.Println("sl3:",l)*/

	// 可以在一行代码中声明并初始化一个 slice 变量。
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)
}

// Slice 是 Go 中一个关键的数据类型，是一个比数组更加强大的序列接口
// Slice 可以组成多维数据结构。内部的 slice 长度可以不同，这和多位数组不同。
// 现在，我们已经看过了数组和 slice，接下来我们将看看 Go 中的另一个关键的内建数据类型：map。
