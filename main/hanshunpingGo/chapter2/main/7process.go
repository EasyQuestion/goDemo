package main

import "fmt"

func main() {

	// type-switch:判断某个interface变量中实际指向的变量类型
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型是%T", i)
	case int:
		fmt.Println("x是int类型")
	case float64:
		fmt.Println("x是float64类型")
	case func(int) float64:
		fmt.Println("x是func(int)类型")
	case bool, string:
		fmt.Println("x是bool或string类型")
	default:
		fmt.Println("未知型")
	}

	/*num := 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough
	case 20:
		fmt.Println("ok2")
		fallthrough
	case 30:
		fmt.Println("ok3")
		fallthrough
	default:
		fmt.Println("没有匹配到...")
	}*/

	/*switch grade := 90; {
	case grade > 90:
		fmt.Println("成绩优秀")
	case grade >= 70:
		fmt.Println("成绩优良")
	case grade >= 69:
		fmt.Println("成绩及格")
	default:
		fmt.Println("不及格...")
	}*/

	/*var age int = 10
	switch {
	case age == 10:
		fmt.Println("age == 10")
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("没有匹配到")
	}*/

	/*var n1, n2, n3 int32 = 5, 20, 5
	switch n1 {
	case 10, n2:
		fmt.Println("ok1")
	case n3:
		fmt.Println("ok2")
	default:
		fmt.Println("没有匹配...")
	}*/
}

/*func main() {
	var myChar byte
	fmt.Println("请输入一个字符：a,b,c,d,e,f,g")
	fmt.Scanf("%c", &myChar)

	switch myChar {
	case 'a':
		fmt.Println("星期一")
	case 'b':
		fmt.Println("星期一")
	case 'c':
		fmt.Println("星期一")
	case 'd':
		fmt.Println("星期一")
	case 'e':
		fmt.Println("星期一")
	case 'f':
		fmt.Println("星期一")
	case 'g':
		fmt.Println("星期一")
	default:
		fmt.Println("其他...")
	}
}*/

/*func main() {

//if age := 20; age > 18 { // 判断语句中允许声明一个变量，变量的作用域只在该条件逻辑块内，其他地方不起作用
//	fmt.Println("你年龄大于18，要对自己的行为负责！")
//}

// 输入人的年龄，如果大于18岁，则提示”你年龄大于18，要对自己的行为负责！“
var age byte
fmt.Println("请输入您的年龄：")
fmt.Scanln(&age)
if age > 18 { // 花括号不能省略，即使只有一行代码
fmt.Println("你年龄大于18，要对自己的行为负责！")
} else {
fmt.Println("你的年龄不大这次放过你了")
}
}*/

// 程序流程控制：顺序、分支、循环
// 向前引用：变量必须先声明，再使用

// if 分支：判断语句中允许声明一个变量，变量的作用域只在该条件逻辑块内，其他地方不起作用
// switch 分支：每个case分支匹配后不需要再加break,只执行一个分支
//            case分支中可以有多个表达式，用逗号间隔,如果case分支是常量或字面量则不能重复（变量就可以，可以骗过编译器）
//            case的值只要和switch表达式的值匹配即可，可以是变量、常量、表达式、函数等等
//            default分支不是必须的
//            特殊用法：switch后没有表达式，相当于多个if-else分支结构，甚至新定义变量
//            switch穿透fallthrough:在case语句块后添加fallthrough，则继续执行下一个case分支的语句块，不用判断
//            type-switch:判断某个interface变量中实际指向的变量类型
