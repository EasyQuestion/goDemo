package model

import (
	"fmt"
)

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (c Customer) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email)
}

func (c *Customer) Copy(target Customer) {
	c.Name = target.Name
	c.Gender = target.Gender
	c.Age = target.Age
	c.Phone = target.Phone
	c.Email = target.Email
}

/*func main() {
	name := "-1"
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name) // 如果这时直接键入enter键，相当于没有赋值，name的值没有变化
	fmt.Println(name == "-1")// true
}*/
