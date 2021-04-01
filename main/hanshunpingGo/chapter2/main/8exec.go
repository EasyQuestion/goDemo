package main

import "fmt"

/*func main() {
	// 声明2个int32型变量并赋值，判断两数之和，如果大于等于50，打印“hello world!”
	var n1, n2 int32 = 10, 50
	if n1+n2 > 50 {
		fmt.Println("hello world!")
	}
}*/

/*func main() {
	// 声明2个float64型的变量并赋值，判断第一个数大于10.0，且第二个数小于20.0，打印两数之和
	var n1, n2 float64 = 11.0, 17.0
	if n1 > 10.0 && n2 < 20.0 {
		fmt.Println(n1, "+", n2, "=", n1+n2)
	}
}*/
/*
func main() {
	// 定义2个变量int32,判断二者的和，是否能被3又能被5整除，打印提示信息
	var n1, n2 int32 = 10, 20
	if sum := n1 + n2; sum%3 == 0 && sum%5 == 0 {
		fmt.Println("两数之和能被3又能被5整除！")
	}
}*/

/*func main() {
	// 判断一个年份是否是闰年（闰年判断条件：方一：能被4整除，但不能被100整除 方二：能被400整除）
	var year int
	fmt.Println("请输入年份：")
	fmt.Scanln(&year)
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("是闰年！")
	} else {
		fmt.Println("不是闰年！")
	}
}*/

//--------------------------------------------------------------------------------------------------
/*func main() {
	// 用switch把小写类型的a,b,c,d,e转成大写，其它的显示Other
	var myChar byte
	fmt.Println("请输入字符：a,b,c,d,e")
	fmt.Scanf("%c", &myChar)

	switch myChar {
	case 'a', 'b', 'c', 'd', 'e':
		fmt.Printf("%c", myChar - 32)
	default:
		fmt.Println("other")
	}
}*/

/*func main() {
	//成绩大于60分，输出“合格”，低于60分，输出“不合格”，输入的成绩不允许大于100
	var score float64
	fmt.Println("请输入成绩：（不能大于100）")
	//fmt.Scanf("%f", &score)
	fmt.Scanln(&score)

	//switch {
	//case score > 100:
	//	fmt.Println("输入的成绩不允许大于100")
	//case score > 60:
	//	fmt.Println("合格")
	//case score < 60:
	//	fmt.Println("不合格")
	//}

	switch int(score / 60) {
	case 1:
		fmt.Println("合格")
	case 0:
		fmt.Println("不合格")
	default:
		fmt.Println("输入有误...")
	}
}*/

/*func main() {
	// 根据用户输入的月份，输出对应的季节，3-5春季，6-8夏季，9-11秋季，12-2冬季
	var month int
	fmt.Println("请输入月份：")
	fmt.Scanln(&month)

	switch month / 3 {
	case 1:
		fmt.Println("春季")
	case 2:
		fmt.Println("夏季")
	case 3:
		fmt.Println("秋季")
	case 0, 4:
		fmt.Println("冬季")
	default:
		fmt.Println("输入有误...")
	}
}*/

func main() {
	//根据用户输入，显示对应内容 星期一：干煸豆角 星期二：醋溜土豆 星期三：红烧狮子头 星期四：油炸花生米 星期五：蒜蓉扇贝 星期六：东北乱炖 星期日：大盘鸡
	var content string
	fmt.Println("请输入：")
	fmt.Scanln(&content)
	switch content {
	case "星期一":
		fmt.Println("干煸豆角")
	case "星期二":
		fmt.Println("醋溜土豆")
	case "星期三":
		fmt.Println("红烧狮子头")
	case "星期四":
		fmt.Println("油炸花生米")
	case "星期五":
		fmt.Println("蒜蓉扇贝")
	case "星期六":
		fmt.Println("东北乱炖")
	case "星期日":
		fmt.Println("大盘鸡")
	default:
		fmt.Println("输入有误...")
	}
}
