package main

import (
	"fmt"
	"goDemo/main/hanshunpingGo/chapter4/model"
)

func main() {
	stu := model.NewStudent("jack", 90.5)
	fmt.Println(*stu)
	// 还是只有首字母大写的才能外包访问
	// 或者通过方法首字母大写来获取对应的值
	fmt.Println("name=", stu.Name, "score=", stu.GetScore())
	stu2 := model.NewStudent2("mary", 100)
	fmt.Println(stu2)
}

// 工厂模式
// golang的结构体中没有构造函数，可以用工厂模式来解决这个问题
// 如果包中的结构体，首字母是小写的，其他包不能访问，又想在其他包创建结构体实例 ---> 工厂模式来解决这个问题
