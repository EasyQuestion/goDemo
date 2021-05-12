package main

import "fmt"

func main() {
	// for(i:=0,j:=100;i<j; i++,j--){  这样的写法不行，必须用平行赋值的方式进行
	for i, j := 0, 100; i < j; i, j = i+1, j-1 {
		fmt.Println(i, j)
	}
}

//func main() {

/*for i := 1; i <= 10; i++ {
	fmt.Println("你好，go语言", i)
}*/

// for的第二种写法
/*i := 1
for i <= 10 {
	fmt.Println("你好，go语言", i)
	i++
}*/

//for的第三种写法
/*i := 1
for { // 相当于 for ;;{
	if i <= 10 {
		fmt.Println("你好，go语言", i)
	} else {
		break
	}
	i++
}*/

//字符串遍历的传统方式
//str := "hello,world"
//str := "hello,world，你好"
/*strArr := []rune(str)
for i := 0; i < len(strArr); i++ {
	fmt.Printf("%c ", strArr[i])
}
fmt.Println()*/

//字符串遍历的range方式
/*for i, v := range str {
	fmt.Printf("i=%d,v=%c\n", i, v)
}*/
//}
//--------------------------------------------
/*func main() {
	// 随机生成1-100之间的一个数，当生成数为99时，输出生成了几次
	var num int
	count := 0
	for {
		count++
		// 为了保证每次输出的值不一样，要设置种子的值不同
		rand.Seed(time.Now().UnixNano())
		num = rand.Intn(100) + 1
		if num == 99 {
			fmt.Println("use times:", count)
			break
		}

	}

}*/
//-------------------------------------------------------
/*func main() {
label2:
	for i := 0; i < 4; i++ {
		//label1:
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break
				//break label1 // 按标签的位置来跳出语句块
				break label2
			}
			fmt.Println("j=", j)
		}
	}
}*/
//-------------------------------------------------------
/*func main() {
	n := 10 // 30
	fmt.Println("ok1")
	if n > 20 {
		goto label1
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	fmt.Println("ok5")
label1:
	fmt.Println("ok6")
	fmt.Println("ok7")
}*/

// for循环
// 基本语法
/*
for 变量初始化;循环条件;变量值迭代 {
	操作语句（也叫循环体）
}
*/
// for range 可以进行字符串或数组的遍历
// 如果字符串中有汉字，那么传统遍历会出现乱码，因为是按字节遍历的，而汉字用utf8编码，一个汉字占3个字节，导致出错
// 解决方式：将string转换成切片数组[]rune
// 而for range是按照字符的方式进行的遍历，所以不会出问题

// 多重循环
// 分析问题时的2大绝招：化难为简，先写死后优化

// break  跳出当前for循环，或中断switch语句
// break + 标签 ：当出现在多层语句块时，可以通过标签来指明跳出的是哪一层语句块

// continue 结束本次循环，继续下一次循环（for,switch）
// continue + 标签 :当出现在多层语句块时，可以通过标签来指明要跳过的是哪一层语句块

// goto 无条件地转移到指定的行继续执行 （go语言中不推荐使用，以免造成程序流程混乱，理解和调试困难）
// goto + 标签
// 汇编语言用goto比较多，高级语言用的比较少

// return 表示跳出所在的方法或函数

// for循环不支持以逗号为间隔的多个赋值语句,必须使用平行赋值的方式来初始化多个变量
