package service

import "goDemo/main/hanshunpingGo/chapter4/project2/model"

type CustomerService struct {
	customerSlice []model.Customer
	customerNum   int
}

func NewCustomerService() CustomerService {
	customerService := CustomerService{customerNum: 1}
	customerService.customerSlice = []model.Customer{
		model.NewCustomer(1, "tom", "男", 42, "112", "tom@sina.com"),
	}
	return customerService
}

func (this *CustomerService) List() []model.Customer {
	return this.customerSlice
}

func (this *CustomerService) Add(name string, gender string, age int, phone string, email string) bool {
	this.customerNum++
	id := this.customerNum
	customer := model.NewCustomer(id, name, gender, age, phone, email)
	this.customerSlice = append(this.customerSlice, customer)
	return true
}

func (this *CustomerService) Delete(id int) bool {
	index := this.findIndexById(id)
	if index == -1 {
		return false
	}
	this.customerSlice = append(this.customerSlice[:index], this.customerSlice[index+1:]...)
	return true
}

func (this *CustomerService) Update(customerUpdate model.Customer) bool {
	index := this.findIndexById(customerUpdate.Id)
	if index == -1 {
		return false
	}
	customer := &this.customerSlice[index]
	customer.Copy(customerUpdate)
	return true
}

func (this *CustomerService) FindCustomerById(id int) *model.Customer {
	index := this.findIndexById(id)
	if index == -1 {
		return nil
	} else {
		return &this.customerSlice[index]
	}
}

func (this *CustomerService) findIndexById(id int) int {
	for i, v := range this.customerSlice {
		if v.Id == id {
			return i
		}
	}
	return -1
}
