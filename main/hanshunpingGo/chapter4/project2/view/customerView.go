package main

import (
	"fmt"
	"goDemo/main/hanshunpingGo/chapter4/project2/model"
	"goDemo/main/hanshunpingGo/chapter4/project2/service"
	"strconv"
)

type customerView struct {
	loop            bool
	customerService service.CustomerService
}

func main() {
	cv := customerView{
		loop: true,
	}
	cv.customerService = service.NewCustomerService()
	cv.startMenu()
}

func (this *customerView) startMenu() {
	choice := 0
	for {
		fmt.Println("------------------客户信息管理软件--------------------")
		fmt.Println("------------------1 添加客户")
		fmt.Println("------------------2 修改客户")
		fmt.Println("------------------3 删除客户")
		fmt.Println("------------------4 客户列表")
		fmt.Println("------------------5 退   出")
		fmt.Println("请选择：（1-5）")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			this.add()
		case 2:
			this.update()
		case 3:
			this.delete()
		case 4:
			this.list()
		case 5:
			this.exit()
		default:
			fmt.Println("输入无效...")
		}

		if !this.loop {
			break
		}
	}
	fmt.Println("退出客户信息管理软件...")
}

func (this *customerView) exit() {
	fmt.Println("确认退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" || choice == "N" || choice == "n" {
			break
		}
		fmt.Println("输入有误，请重新输入：y/n")
	}
	if choice == "Y" || choice == "y" {
		this.loop = false
	}
}

func (this *customerView) list() {
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, v := range this.customerService.List() {
		fmt.Println(v.GetInfo())
	}
	fmt.Println()
}

func (this *customerView) add() {
	fmt.Println("请输入姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("请输入性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("请输入年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("请输入电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("请输入邮箱：")
	email := ""
	fmt.Scanln(&email)
	if this.customerService.Add(name, gender, age, phone, email) {
		fmt.Println("---------------添加成功-------------------")
	} else {
		fmt.Println("---------------添加失败-------------------")
	}
	fmt.Println()
}

func (this *customerView) delete() {
	id := -1
	fmt.Print("请选择待删除客户编号（-1退出）：")
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	if this.customerService.Delete(id) {
		fmt.Println("-----------删除完成-------------")
	} else {
		fmt.Println("-----------删除失败-------------")
	}
	fmt.Println()
}

func (this *customerView) update() {
	id := -1
	fmt.Print("请选择待修改客户编号（-1退出）：")
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	customer := this.customerService.FindCustomerById(id)
	if customer == nil {
		fmt.Println("编号", id, "的客户不存在！")
		return
	}
	customerNew := model.Customer{}
	customerNew.Id = id
	customerNew.Copy(*customer)
	fmt.Println("姓名（", customer.Name, "）:")
	name := ""
	fmt.Scanln(&name)
	if name != "" {
		customerNew.Name = name
	}
	fmt.Println("性别（", customer.Gender, "）：")
	gender := ""
	fmt.Scanln(&gender)
	if gender != "" {
		customerNew.Gender = gender
	}
	fmt.Println("年龄（", customer.Age, "）：")
	age := ""
	fmt.Scanln(&age)
	if age != "" {
		ageInt, err := strconv.Atoi(age)
		for err != nil {
			fmt.Println("年龄数据有误...请重新输入（-1退出修改）：")
			fmt.Scanln(&age)
			if age == "-1" {
				return
			}
			ageInt, err = strconv.Atoi(age)
		}
		customerNew.Age = ageInt
	}
	fmt.Println("电话（", customer.Phone, "）：")
	phone := ""
	fmt.Scanln(&phone)
	if phone != "" {
		customerNew.Phone = phone
	}
	fmt.Println("邮箱（", customer.Email, "）：")
	email := ""
	fmt.Scanln(&email)
	if email != "" {
		customerNew.Email = email
	}

	if this.customerService.Update(customerNew) {
		fmt.Println("-----------修改完成-------------")
		return
	} else {
		fmt.Println("-----------修改失败-------------编号", id, "的客户不存在！")
		return
	}
}
