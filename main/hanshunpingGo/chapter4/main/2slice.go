package main

import "fmt"

/*func main() {

	intArr := [...]int{1, 22, 33, 66, 99}
	// 切片slice 引用数组 intArr 下标[1-3) 左闭右开的 元素
	slice := intArr[1:3]
	intArr[1] = 24
	fmt.Println("slice=", slice)
	fmt.Println("slice len=", len(slice))      //2 slice的长度
	fmt.Println("slice capacity=", cap(slice)) //4 slice的容量：可以动态变化，一般是元素个数的2倍或者1.5倍

	fmt.Printf("intArr[1]的地址是:%p\n", &intArr[1])
	fmt.Printf("slice[0]的地址是:%p\n", &slice[0])
}*/

/*func main() {
	// slice使用的3种方式：
	// 方一：
	intArr := [...]int{1, 22, 33, 66, 99}
	slice := intArr[1:3]
	//slice[4] = 12 // 编译通过，但是运行报错 index out of range
	slice = append(slice, 77) // 会将引用的数组内的元素修改
	fmt.Println("intArr=", intArr)
	fmt.Println("slice=", slice, "len=", len(slice), "cap=", cap(slice))
	// 切片还可以再切片
	sliceOther := slice[1:]
	fmt.Println("sliceOther = ", sliceOther, "len=", len(sliceOther), "cap=", cap(sliceOther))
	sliceOther[0] = 100
	fmt.Println("intArr = ", intArr)

	// 方二：make
	slice = make([]int, 4, 10) // 定义slice的len,cap,cap是可选的（如果cap写了，必须要>=len）
	fmt.Printf("slice的len:%d,slice的cap:%d\n", len(slice), cap(slice))

	var slice2 []float64
	fmt.Println("slice2=", slice2)
	slice2 = make([]float64, 4, 10)
	fmt.Println("slice2=", slice2)

	// 方三：直接初始化
	slice3 := []string{"tom", "jack", "mary"}
	fmt.Printf("slice3:%v len:%d cap:%d", slice3, len(slice3), cap(slice3))

}*/

/*func main() {
	slice := []int{100, 200, 300}
	// 对append()的结果，必须重新赋值，才能达到扩容或修改的效果
	// append()可以同时追加多个值
	slice = append(slice, 400, 500, 600)
	// append()也可以追加一个切片,语法是 sliceNew...
	slice = append(slice, slice...)
	slice[0] = 99
	fmt.Println("slice=", slice)
}*/

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 1)
	fmt.Println("slice2=", slice2)
	// slice 和 slice2的数据空间是相互独立的，互不影响
	copy(slice2, slice)
	slice[0] = 10
	fmt.Println("slice=", slice)
	fmt.Println("slice2=", slice2)
}

// slice切片：个数可变
// 切片是数组的一个引用，因此切片是引用类型，在进行传递时遵循引用传递的机制
// 切片的使用和数组类似：遍历切片、访问切片元素、求切片长度
// 切片长度是可变的，因此切片是一个可以动态变化的数组
// 切片定义基本语法 var a []int

// slice的内存布局：可以认为基本是3部分：首地址（指针类型）  |  len  |  cap
// 从底层来说，slice本质是struct结构体类型
/* type slice struct {
	ptr *[2]int // 指针的具体类型要看具体的定义
    len int
    cap int
 }
*/

// slice使用的3种方式：
// slice := intArr[1:3]  intArr[1:]  intArr[:4] intArr[:]
// slice := make([]int,4,10)
//         make生成的切片，所引用的数组只能由该切片访问，其它变量不能访问
// slice := []int{1,3,5} 此时 len(slice) == cap(slice)
//         相当于make一个切片

// slice遍历：和数组一样，for传统遍历，for-range遍历

// slice的使用注意事项：
// 1.slice := arr[1:3] 截取方式创建切片时，注意下标位置是左闭右开
// 2.切片初始化时不能越界，范围在[0-len(arr)]之间，但是可以动态增长
// 3.cap() 为go内置函数，统计切片的容量
// 4.切片定义完之后，还不能使用，必须引用一个数组或make一个空间供切片引用
// 5.切片可以继续切片
// 6.append()内置函数，对切片进行动态增长，
//                会对切片进行扩容，可能会生成新的基本数组，必须接收它的结果
// 7.copy()内置函数，对切片进行拷贝
//                只有切片才能copy()拷贝，不能是数组
//                拷贝时，2个切片的数据空间是相互独立的，互不影响
//                2个切片的长度可以不一致
// 8.切片是引用类型，在进行传递时遵循引用传递的机制
