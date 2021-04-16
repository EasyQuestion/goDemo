package main

import "fmt"

func main() {
	/*
		0 0 0 0 0 0
		0 0 1 0 0 0
		0 2 0 3 0 0
		0 0 0 0 0 0
	*/
	// 二维数组使用方式一：先定义/声明，再赋值
	var arr [4][6]int
	fmt.Println("arr=", arr)
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	fmt.Println("arr=", arr)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
	// 二维数组使用方式二：直接初始化
	arr3 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr3=", arr3)

	arr4 := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr4=", arr4)

	// 二维数组在内存中的布局
	// 相当于存储了一组指针，每个指针对应一个一维数组的首地址
	var arr2 [2][1]int
	fmt.Printf("arr2[0]的地址%p\n", &arr2[0])
	fmt.Printf("arr2[1]的地址%p\n", &arr2[1])
	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][0])
	fmt.Printf("arr2[1][0]的地址%p\n", &arr2[1][0])
}

// 多维数组 -> 简化：二维数组
// 二维数组的2种使用方式：先定义/声明，再赋值；直接初始化
// 二维数组在内存中的布局
// 二维数组的遍历：传统for方式、for-range方式
