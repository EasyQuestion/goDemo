package main

import "fmt"

// 冒泡排序法
//func BubbleSort(arr []int) {
/*24, 69, 80, 57, 13
第一次外层循环
	第1次比较 结果：24, 69, 80, 57, 13
	第2次比较 结果：24, 69, 80, 57, 13
	第3次比较 结果：24, 69, 57, 80, 13
	第4次比较 结果：24, 69, 57, 13, 【80】
第2次外层循环
	第1次比较 结果：24, 69, 57, 13, 【80】
	第2次比较 结果：24, 57, 69, 13, 【80】
	第3次比较 结果：24, 57, 13, 【69】, 【80】
第3次外层循环
	第1次比较 结果：24, 57, 13, 【69】, 【80】
	第2次比较 结果：24, 13, 【57】, 【69】, 【80】
第4次外层循环
	第1次比较 结果：13, 【24】, 【57】, 【69】, 【80】
*/
// 思路：先写内层循环，再写外层循环
// 先实现第一次外层循环的4次比较
/*for j := 0; j < len(arr)-1; j++ {
		for i := 0; i < len(arr)-j-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

func main() {
	// 用冒泡排序法将数组排列成一个从小到大的有序数列
	arr := []int{24, 69, 80, 57, 13}
	BubbleSort(arr)
	fmt.Println("arr=", arr)
}*/
//--------------------------------------------
/*func main() {
	// 顺序查找：从键盘输入一个名字，判断已有数列中是否含有这个名字
	nameArr := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var name string
	fmt.Println("请输入一个名字：")
	fmt.Scanln(&name)

	// 第一种方式：
	//for i, n := range nameArr {
	//	if strings.Contains(n, name) {
	//		fmt.Println(n, "包含这个名字", name)
	//		break
	//	} else if i == len(nameArr)-1 {
	//		fmt.Println("没有找到...")
	//	}
	//}

	// 第二种方式： 推荐使用第二种方式（更简洁清晰）
	index := -1
	for i, n := range nameArr {
		if strings.Contains(n, name) {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Println(nameArr[index], "包含这个名字", name)
	} else {
		fmt.Println("没有找到...")
	}
}*/
//--------------------------------------------
func findByBinary(leftIndex, rightIndex int, arr []int, value int) int {

	middleIndex := (leftIndex + rightIndex) / 2
	if arr[middleIndex] == value {
		return middleIndex
	} else if leftIndex == rightIndex {
		return -1
	} else if arr[middleIndex] > value {
		return findByBinary(leftIndex, middleIndex-1, arr, value)
	} else if arr[middleIndex] < value {
		return findByBinary(middleIndex+1, rightIndex, arr, value)
	}
	return -1
}

func main() {
	// 对一个有序数组进行二分查找 {1，8，10，89，1000，1234}
	// 可以使用递归
	arr := []int{1, 8, 10, 89, 1000, 1234}
	var n int
	fmt.Println("请输入要查找的数字：")
	fmt.Scanln(&n)
	index := findByBinary(0, len(arr)-1, arr, n)
	if index != -1 {
		fmt.Println("index=", index)
	} else {
		fmt.Println("没有找到...")
	}
}

// 排序和查找
// 排序：内部排序和外部排序
// 内部排序：将需要处理的所有数据都加载到内存中，进行排序（交换排序法、选择排序法、插入排序法）
// 外部排序：数据量过大，无法全部加载到内存中，需要借助外部存储进行排序（合并排序法、直接合并排序法）

// 交换排序法：冒泡排序法Bubble sort、快速排序法Quick sort
// 重点讲冒泡排序法
// 思路：先易后难，先写内层，再写外层

// 查找（常用）：顺序查找、二分查找
// 二分查找，前提：该数组是有序的
