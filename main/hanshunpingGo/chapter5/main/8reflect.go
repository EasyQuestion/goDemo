package main

import (
	"fmt"
	"reflect"
)

func test1(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Printf("rTyp=%v rTyp type=%T\n", rTyp, rTyp)

	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)
	// reflect获取的value值并不能直接使用该类型，需要用 rVal.Int() 转换后才能正常使用
	num := 2 + rVal.Int()
	fmt.Println("num=", num)
	// 传入的是int类型，这里用float类型来处理，编译通过，但是运行时会报错
	// panic: reflect: call of reflect.Value.Float on int Value
	//num2 := rVal.Float()
	//fmt.Println("num2=", num2)

	// 如果不是基础类型，需要先转换成interface{}类型，再通过类型断言转换成对应的类型，之后才能正常使用
	// 如果不确定转换后的类型，需要通过switch来进行类型判断
	res := rVal.Interface().(int)
	fmt.Printf("res=%v res type=%T\n", res, res)
}

func test02(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp = ", rTyp)

	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v rVal type = %T\n", rVal, rVal)

	fmt.Println("rTyp kind=", rTyp.Kind(), "rVal kind=", rVal.Kind())

	// 需要强调下，反射只能获取运行时的对象信息，所以必须要先转换才能获取到具体类型的属性值
	// 编译器无法判断出运行时的对象类型
	student, ok := rVal.Interface().(Student)
	if ok {
		fmt.Println("student name=", student.Name)
	}
}

func test03(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal kind=%v\n", rVal.Kind())
	//rVal.SetInt(20) // 直接这样写会报错 panic: reflect: reflect.Value.SetInt using unaddressable value
	// rVal.Elem() 相当于取到了指针指向的地址
	rVal.Elem().SetInt(20)
}

type Student struct {
	Name string
	Age  int
}

func main() {
	num := 10
	//test1(num)
	// 通过反射赋值
	test03(&num)
	fmt.Println("num=", num)

	//stu := Student{
	//	Name: "tom",
	//	Age:  12,
	//}
	//test02(stu)
}

// 反射 reflect包：运行时反射
// type类别  kind类型
// 应用：1.json的序列化和反序列化 2.go框架开发（比如适配器函数）
// reflect中重要的概念 reflect.Type(本质是接口)  reflect.Value(本质是结构体)
//                   reflect.TypeOf()   reflect.ValueOf()
//                   变量、interface{}、reflect.Value三者的相互转换

// 反射的使用注意事项
// 1.type类别(具体哪个类)  kind类型(类型常量) 相当于  kind->电器，type->冰箱
//   获取kind的2种方式：1-rTyp.Kind() 2-rVal.Kind()
// 2.kind,type有可能相同，也有可能不同
// 3.变量、interface{}、reflect.Value三者的相互转换
//  变量a 通过方法入参匹配方法，转换成 interface{} b  通过  c=reflect.ValueOf(b)  方法，转换成 reflect.Value
//  reflect.Value 通过 c.interface()  方法，转换成 interface{} b，通过  类型断言 转换成 变量a
// 4.使用反射方式获取变量的值，要求数据类型匹配，否则报错panic
// 5.Elem()  获取对指针类型的interface{}具体的值
