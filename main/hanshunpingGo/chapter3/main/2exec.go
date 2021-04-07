package main

import "fmt"

/*func Fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fibonacci(n-2) + Fibonacci(n-1)
	}
}

func main() {
	// 斐波那契数 1，1，2，3，5，8，13...求出第n个数的值是多少
	num := Fibonacci(7)
	fmt.Println("num = ", num)
}*/
/*
func funcN(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*funcN(n-1) + 1
	}
}
func main() {
	// f(1) = 3;f(n) = 2 * f(n-1)+1;求出f(n)的值
	fmt.Println(funcN(1))
	fmt.Println(funcN(2))
	fmt.Println(funcN(3))
	fmt.Println(funcN(4))
}*/

/* 猴子吃桃子问题：
猴子第一天吃了若干个桃子，当即吃了一半，还不解馋，又多吃了一个；
第二天，吃剩下的桃子的一半，还不过瘾，又多吃了一个；
以后每天都吃前一天剩下的一半多一个，
到第10天想再吃时，只剩下一个桃子了。
问第一天共吃了多少个桃子？
*/

func monkeyPeach(days int, num int) int {
	if days > 1 {
		days--
		num = (num + 1) * 2
		return monkeyPeach(days, num)
	} else {
		return num
	}
}

func sum(n1, n2 float32) float32 { // 这样的写法也可以
	fmt.Printf("n1 type=%T\n", n1)
	return n1 + n2
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {

	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println("a=", a, "b=", b)

	//fmt.Println("sum=", sum(1, 2))

	//fmt.Println(monkeyPeach(10, 1))
	// 10天：1
	// 9天：（1+1）*2 = 4
	// 8天：（4+1）*2= 10
	// 7天：22
	// 6天：46
	// 5天：94
	// 4天：190
	// 3天：382
	// 2天：766
	// 1天：1534
	/*num := 1
	for i := 1; i < 10; i++ {
		num = (num + 1) * 2
	}
	fmt.Println("num=", num)*/
}
