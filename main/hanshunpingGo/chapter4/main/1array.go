package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

	var arr []int
	fmt.Printf("arr的类型：%T ", arr) // []int
}*/

/*func main() {
	heroes := [3]string{"宋江", "吴用", "卢俊义"}
	for i, v := range heroes {
		fmt.Printf("i=%v v=%v", i, v)
	}
	fmt.Println()
	for _, v := range heroes {
		fmt.Printf("v=%v", v)
	}
}*/

/*func main() {
	var byteArr [3]byte
	fmt.Println(byteArr) //[0 0 0]
}*/

//----------------------------------------
/*func test(arr [3]int) {
	arr[0] = 12
	arr[1] = 13
	arr[2] = 14
	fmt.Println("test:arr-", arr)
}
func test2(arr *[3]int) {
	arr[0] = 12 // go会自动解引用，正规的写法是 (*arr)[0] = 12
	arr[1] = 13
	arr[2] = 14
	fmt.Println("test2:arr-", *arr)
}

// 数组属于值类型，在默认情况下会进行值传递，因此会进行值拷贝，数组间不会相互影响
func main() {
	arr := [3]int{1, 2, 3}
	test(arr)
	fmt.Println(arr)
	test2(&arr)
	fmt.Println(arr)
}*/
//-----------------------------------------------
/*func main() {
	// 打印A-Z
	var myChars []byte
	for i := 0; i < 26; i++ {
		myChars = append(myChars, 'A'+byte(i))
	}
	fmt.Println(string(myChars))
}*/

/*func main() {
	// 找到数组中的最大值，并输出下标
	var maxIndex int
	intArr := [5]int{3, 100, 23, 97, 55}
	for i, v := range intArr {
		if v > intArr[maxIndex] {
			maxIndex = i
		}
	}
	fmt.Println("maxIndex=", maxIndex)
}*/
/*func main() {
	// 求出一个数组的和，平均值
	intArr := [...]int{3, 100, 23, 97, 55}
	totalValue := 0
	for _, v := range intArr {
		totalValue += v
	}
	avgValue := float64(totalValue) / float64(len(intArr))
	fmt.Printf("totalValue=%v avgValue=%.2f", totalValue, avgValue)
}*/
//--------------------------------------------------------
func main() {
	//随机生成5个数，并将其反转打印
	var intArr [5]int
	arrLen := len(intArr)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < arrLen; i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Println("arr:", intArr)
	for i := 0; i < arrLen/2; i++ {
		intArr[i], intArr[arrLen-1-i] = intArr[arrLen-1-i], intArr[i]
	}
	fmt.Println("reverse arr:", intArr)
}

// 数组：可以存放多个同一类型的数据
// 在go中，数组是值类型
// 内存布局：连续的内存区域，数组的地址就是第一个元素的地址
// 初始化数组的4种方式：[3]int , [...]int{1:80,0:90,2:99} ,
// 遍历数组 for-range

// 数组使用的注意事项
// 1.数组是多个相同类型的数据组合，一个数组一旦声明/定义了，其长度是固定的，不能变动
//                             （[3]int [4]int是两种数据类型，因为长度不同）
// 2.var arr []int --> 这是slice切片的定义，后面会讲
// 3.数组中的元素可以是任何数据类型，包括值类型和引用类型，但是不能混用
// 4.数组创建后，如果没有赋值，有默认值
// 5.数组的使用步骤 1）声明数组并开辟空间 2）给数组各个元素赋值(或者用默认值) 3）使用数组
// 6.数组下标从0开始
// 7.数组下标必须在指定范围内，否则报panic,数组越界，比如var arr [5]int的有效下标为0-4
// 8.go中的数组属于值类型，在默认情况下会进行值传递，因此会进行值拷贝，数组间不会相互影响
// 9.如果想在其它函数中，修改原数组的值，可以使用指针进行引用传递
// 10.go中长度是数组的一部分，在传递函数参数时，需要考虑数组的长度
