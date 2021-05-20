package main

import (
	"fmt"
	"reflect"
)

type ReflectStudent struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score int
}

func (stu ReflectStudent) Print() {
	fmt.Println("print start...")
	fmt.Println(stu)
	fmt.Println("print end...")
}

func (stu ReflectStudent) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (stu ReflectStudent) SetValue(name string, age int, score int) {
	stu.Name = name
	stu.Age = age
	stu.Score = score
}

func reflectStruct(stu ReflectStudent) {
	typ := reflect.TypeOf(stu)
	val := reflect.ValueOf(stu)
	kind := val.Kind()
	if kind != reflect.Struct {
		fmt.Println("type is not struct ,exit...")
		return
	}
	fieldCount := val.NumField()
	fmt.Println("fieldCount=", fieldCount)
	for i := 0; i < fieldCount; i++ {
		fmt.Print("field", i, " ", typ.Field(i).Name, " ", typ.Field(i).Type, ":", val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Print(" json:", tagVal)
		}
		fmt.Println()
	}

	methodCount := val.NumMethod()
	fmt.Println("methodCount=", methodCount)
	val.Method(1).Call(nil)

	params := []reflect.Value{
		reflect.ValueOf(10), reflect.ValueOf(20),
	}
	result := val.Method(0).Call(params)
	fmt.Println("result=", result[0])
}

func main() {
	// 通过反射获取结构体的字段、调用结构体的方法，并获取结构体标签的值
	stu := ReflectStudent{
		Name:  "陈济恒",
		Age:   300,
		Score: 100,
	}
	reflectStruct(stu)
}

// reflect.Value.Field(i) 拿到结构体的字段
// reflect.Value.NumField() 结构体字段个数

// reflect.Type.Field(i).Tag.Get(xxx)字段对应的tag标签

// reflect.Value.Method(i) 拿到结构体的方法
// reflect.Value.NumMethod() 拿到结构体方法个数

// reflect.Value.Method.Call(xxx)调用方法
//            注：方法的排序，默认按照方法名的ASCII码排序
//                Call()方法，传入的是[]reflect.Value,返回的也是[]reflect.Value

// 通过反射修改对象的值 reflect.Value.FieldByName(xxx).setString(xxx)
// 适配器函数：通过反射来调用函数 reflect.Method(i).Call(xxx)
// 通过反射来新建结构体对象 reflect.New(xxx).Elem()  注意一定要用Elem()，reflect.New(xxx)返回的是指针类型