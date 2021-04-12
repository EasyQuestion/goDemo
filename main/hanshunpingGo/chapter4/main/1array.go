package main

import "fmt"

/*func main() {

	// 入门案例
	var hens [6]float64
	hens[0] = 3
	hens[1] = 5
	hens[2] = 1
	hens[3] = 3.4
	hens[4] = 2
	hens[5] = 50

	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}
	avgWeight := fmt.Sprintf("%.2f", totalWeight/float64(len(hens)))
	fmt.Println("totalWeight=", totalWeight, "avgWeight=", avgWeight)
}*/

/*func main() {
	var intArr [3]int
	fmt.Println(intArr)
	fmt.Printf("intArr的地址：%p\n", &intArr)
	fmt.Printf("intArr[0]的地址：%p\n", &intArr[0])
	fmt.Printf("intArr[1]的地址：%p\n", &intArr[1])
	fmt.Printf("intArr[2]的地址：%p\n", &intArr[2])
}*/

/*func main() {
	//初始化数组的4种方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Printf("numArr01的类型：%T ", numArr01)
	fmt.Println("numArr01=", numArr01)

	var numArr02 = [3]int{5, 6, 7}
	fmt.Printf("numArr02的类型：%T ", numArr02)
	fmt.Println("numArr02=", numArr02)

	var numArr03 = [...]int{8, 9, 10} // 这里的[...]是固定写法
	fmt.Printf("numArr03的类型：%T ", numArr03)
	fmt.Println("numArr03=", numArr03)

	var numArr04 = [...]int{1: 800, 0: 900, 2: 999}
	fmt.Printf("numArr04的类型：%T ", numArr04)
	fmt.Println("numArr04=", numArr04)

	strArr05 := [...]string{1: "tom", 0: "jack", 2: "mary"}
	fmt.Printf("strArr05的类型：%T ", strArr05)
	fmt.Println("strArr05=", strArr05)
}*/

func main() {
	heroes := [3]string{"宋江", "吴用", "卢俊义"}
	for i, v := range heroes {
		fmt.Printf("i=%v v=%v", i, v)
	}
	fmt.Println()
	for _, v := range heroes {
		fmt.Printf("v=%v", v)
	}
}

// 数组：可以存放多个同一类型的数据
// 在go中，数组是值类型
// 内存布局：连续的内存区域，数组的地址就是第一个元素的地址
// 初始化数组的4种方式：[3]int , [...]int{1:80,0:90,2:99} ,
// 遍历数组 for-range
