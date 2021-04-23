package familyAccount

import "fmt"

type familyAccount struct {
	choice     string
	loop       bool
	detail     string
	balance    float64
	money      float64
	note       string
	isNoDetail bool
}

func NewAccount() *familyAccount {
	return &familyAccount{
		choice:     "",
		loop:       false,
		detail:     "收支\t账户金额\t收支金额\t收支说明",
		balance:    10000.0,
		money:      0.0,
		note:       "",
		isNoDetail: true,
	}
}

func (this *familyAccount) StartMenu() {
	for {
		fmt.Println("\n---------------家庭记账收支系统-------------------")
		fmt.Println("               1.展示详情")
		fmt.Println("               2.登记收入")
		fmt.Println("               3.登记支出")
		fmt.Println("               4.退出系统")
		fmt.Print("请选择：")
		fmt.Scanln(&this.choice)

		switch this.choice {
		case "1":
			this.showInfo()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("输入有误，请重新选择")
		}

		if this.loop {
			break
		}
	}
	fmt.Println("退出家庭记账收支系统...")
}

func (this *familyAccount) showInfo() {
	fmt.Println("---------------当前收支明细记录-------------------")
	if this.isNoDetail {
		fmt.Println("还没有收支，请记一笔吧..")
	} else {
		fmt.Println(this.detail)
	}
}

func (this *familyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.balance += this.money
	this.detail += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.isNoDetail = false
}

func (this *familyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
		return
	}
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	this.balance -= this.money
	this.detail += fmt.Sprintf("\n支出\t%v\t%v\t%v\n", this.balance, this.money, this.note)
	this.isNoDetail = false
}

func (this *familyAccount) exit() {
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
		this.loop = true
	}
}
