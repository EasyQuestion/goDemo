package main

import "fmt"

func main() {
	a := 9
	b := 2
	a, b = b, a
	fmt.Println("a=", a, "b=", b)
}

//------------------------------------------------
/*func main() {
	i := 10
	//if i > 9 && test() {
	if i < 9 && test() {
		fmt.Println("ok...")
	}
}

func test() bool {
	fmt.Println("test...")
	return true
}*/
//---------------------------------------------------
// 课堂练习
/*func main() {
	// 1-假如还有97天放假，几个星期零几天？
	day := 97
	weeks, days := leastDays(day)
	fmt.Println("还有97天放假，对应", weeks, "个星期零", days, "天")

	// 2-定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：5/9*(华氏温度-100)，求出华氏温度对应的摄氏温度
	f := 134.2
	fmt.Println("华氏温度", f, "对应的摄氏温度是：", fahrenheitToCentigrade(float32(f)))
}

func leastDays(day int) (int, int) {
	weeks := day / 7
	days := day % 7
	return weeks, days
}

func fahrenheitToCentigrade(f float32) float32 {
	return 5.0 / 9 * (f - 100)
}*/
//--------------------------------------------------
/*func main() {
	fmt.Println("10/4=", 10/4)
	fmt.Println("10.0/4=", 10.0/4)

	fmt.Println("10%3=", 10%3)
	fmt.Println("-10%3=", -10%3)
	fmt.Println("10%-3=", 10%-3)
	fmt.Println("-10%-3=", -10%-3)

	// a++,a--只能作单独的语句使用
	i := 10
	i++
	fmt.Println("i=", i)
	i--
	fmt.Println("i=", i)
}*/

// 运算符:
// 1.算术运算符(对数值型变量运算)+ - * / % a++ a-- a+b
//  注意：go中没有 ++a --a,而且a++,a--只能作单独的语句使用，不能这样使用 b := a++
// 2.赋值运算符
//    = += -= *= /= %=
//      (操作顺序是从右向左)(运算符左边只能是变量，右边则可以是变量、表达式、常量值)
//    <<= >>= &= |= ^=(按位异或后赋值)  主要涉及到二进制的赋值运算
// 3.比较运算符（关系运算符）== != > < >= <=
//  (结果都是bool类型的，只有true,false)(经常用在if,for循环中)
// 4.逻辑运算符(连接多个条件，即连接多个条件表达式，最终的结果也是bool类型)  && || ! (也叫短路与、短路或)
// 5.位运算符 & | ^ << >>
// 6.其它运算符 & *
// go中不支持三元运算符，可以用if,else代替


// 运算符的优先级
/*优先级 	运算符
7 		^ !
6 		* / % << >> & &^
5 		+ - | ^
4 		== != < <= >= >
3 		<-
2 		&&
1 		||
*/
