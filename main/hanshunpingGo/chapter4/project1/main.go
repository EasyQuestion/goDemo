package main

//第二版（面向对象）-------------------------------------------

import (
	"fmt"
	"goDemo/main/hanshunpingGo/chapter4/project1/familyAccount"
)

func main() {
	fmt.Println("start family account:")
	account := familyAccount.NewAccount()
	account.StartMenu()
}

//第一版（面向过程）-------------------------------------------
/*
func main() {
	choice := ""
	loop := false
	detail := "收支\t账户金额\t收支金额\t收支说明"
	balance := 10000.0
	money := 0.0
	note := ""
	isNoDetail := true
	for {
		fmt.Println("\n---------------家庭记账收支系统-------------------")
		fmt.Println("               1.展示详情")
		fmt.Println("               2.登记收入")
		fmt.Println("               3.登记支出")
		fmt.Println("               4.退出系统")
		fmt.Print("请选择：")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			fmt.Println("---------------当前收支明细记录-------------------")
			if isNoDetail {
				fmt.Println("还没有收支，请记一笔吧..")
			} else {
				fmt.Println(detail)
			}
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			balance += money
			detail += fmt.Sprintf("\n收入\t%v\t%v\t%v\n", balance, money, note)
			isNoDetail = false
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			balance -= money
			detail += fmt.Sprintf("\n支出\t%v\t%v\t%v\n", balance, money, note)
			isNoDetail = false
		case "4":
			fmt.Println("确认要退出吗？y/n")
			exitEnsure := ""
			for {
				fmt.Scanln(&exitEnsure)
				if exitEnsure == "y" || exitEnsure == "n" {
					break
				}
				fmt.Println("输入有误，请重新选择 y/n")
			}
			if exitEnsure == "y" {
				loop = true
			}
		default:
			fmt.Println("输入有误，请重新选择")
		}

		if loop {
			break
		}
	}
	fmt.Println("退出家庭记账收支系统...")
}
*/
