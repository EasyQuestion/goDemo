package main

import "fmt"

/*func main() {
	// 打印1-100间所有是9的倍数的整数的个数及总和
	var count, sum int
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Println("1-100间9个倍数的整数个数为", count, "它们的总和是", sum)
}*/
/*
func main() {
	// 输入一个大于0的整数，输出所有两数和是它的表达式
	var num int
	fmt.Println("请输入一个大于0的整数：")
	fmt.Scanln(&num)
	for i := 0; i <= num; i++ {
		fmt.Println(i, "+", num-i, "=", num)
	}
}*/

/*func main() {
	// 有3个班的成绩，每班5人，要求输出每个班的平均成绩和3个班的总平均成绩
	score := [3][5]float64{{89, 76, 90, 92, 95}, {100, 98, 78, 89, 90}, {89, 90, 92, 92, 98}}
	var totalAverage, totalSum float64
	for i := 0; i < 3; i++ {
		var average, sum float64
		for j := 0; j < 5; j++ {
			sum += score[i][j]
		}
		average = sum / 5
		totalSum += sum
		fmt.Printf("%d班的平均成绩：%.2f\n", i+1, average)
	}

	totalAverage = totalSum / 3 / 5
	fmt.Printf("3个班的总平均成绩：%.2f", totalAverage)
}*/

/*func main() {
	// 打印空心金字塔
	// 先打印3层的，再改成多层的
	// 扩展：打印菱形
	tail := 20
	for i := 1; i <= tail; i++ {
		printStar(tail, i)
	}
	for i := tail - 1; i > 0; i-- {
		printStar(tail, i)
	}
}

func printStar(tail int, currentFloor int) {
	for k := tail - currentFloor; k > 0; k-- {
		fmt.Print(" ")
	}
	for j := 1; j <= 2*currentFloor-1; j++ {
		if j == 1 || j == 2*currentFloor-1 {
			fmt.Print("*")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
*/

/*func main() {
	// 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, "*", i, "=", j*i, " ")
		}
		fmt.Println()
	}
}*/

func main() {
	// 某人有100000元，每经过一次路口，需要交费，规则如下：
	// 当现金 > 50000时，每次交5%
	// 当现金 <= 50000时，每次交1000
	// 计算出此人可以过几次路口，for break方式完成
	var money float64 = 100000
	var passCount = 0
	for {
		if money <= 0 {
			break
		}
		passCount++
		if money > 50000 {
			money = money - money*0.05
		} else if money <= 50000 {
			money -= 1000
		}
	}
	fmt.Println("此人可以经过", passCount, "次路口")
}
