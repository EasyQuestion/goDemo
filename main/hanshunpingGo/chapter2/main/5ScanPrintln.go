package main

import "fmt"

func main() {

	// 可以从控制台接收用户信息（姓名、年龄、工资，是否通过考试）
	var name string
	var age int
	var salary float64
	var isPassedTest bool
	//fmt.Println("请输入您的姓名：")
	//fmt.Scanln(&name)
	//fmt.Println("请输入您的年龄：")
	//fmt.Scanln(&age)
	//fmt.Println("请输入您的工资：")
	//fmt.Scanln(&salary)
	//fmt.Println("您是否通过考试，请输入true或false:")
	//fmt.Scanln(&isPassedTest)
	//fmt.Println("您的信息为【姓名：", name, "年龄：", age, "工资：", salary, "是否通过考试：", isPassedTest, "】")

	fmt.Println("请输入姓名 年龄 工资 是否通过考试，用空格隔开：")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPassedTest)
	fmt.Println("您的信息为【姓名：", name, "年龄：", age, "工资：", salary, "是否通过考试：", isPassedTest, "】")
}

//在go 语言中，键盘输入用户信息 fmt.Scanln()  fmt.Scanf()
