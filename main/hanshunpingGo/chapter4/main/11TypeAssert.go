package main

import "fmt"

// 总结：当带有类型断言的方法调用时，指针可以自动解引用，但是对象不能自动转换成指针
func main() {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer).Add(b)
	fmt.Println(sum)
}

// 有2种方式，方一：
/*type Integer int
func (a Integer) Add(b Integer) Integer {
	return a + b
}*/

// 有2种方式，方二：
type Integer int
func (a *Integer) Add(b Integer) Integer {
	return *a + b
}

/*func main() {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = a
	sum := i.(Integer).Add(b)
	fmt.Println(sum)
}

type Integer int
func (a Integer) Add(b Integer) Integer {
	return a + b
}*/

//------------------------------------------------------------------
/*type Student struct{}

func TypeJudge(params ...interface{}) {
	for i, v := range params {
		switch v.(type) {
		case int, int32, int64:
			fmt.Printf("第%d个参数是 整数 类型，它的值是 %v\n", i+1, v)
		case float32, float64:
			fmt.Printf("第%d个参数是 float 类型，它的值是 %v\n", i+1, v)
		case bool:
			fmt.Printf("第%d个参数是 bool 类型，它的值是 %v\n", i+1, v)
		case nil:
			fmt.Printf("第%d个参数是 nil 类型，它的值是 %v\n", i+1, v)
		case string:
			fmt.Printf("第%d个参数是 string 类型，它的值是 %v\n", i+1, v)
		case Student:
			fmt.Printf("第%d个参数是 Student 类型，它的值是 %v\n", i+1, v)
		case *Student:
			fmt.Printf("第%d个参数是 *Student 类型，它的值是 %v\n", i+1, v)
		default:
			fmt.Println("第%d个参数 类型不确定，它的值是 %v\n", i+1, v)
		}
	}
}

func main() {
	var n1 float32 = 8.9
	var n2 float64 = 1.1
	var n3 int = 12
	var name string = "tom"
	isPassed := true
	address := "12#3-301"
	stu := Student{}
	TypeJudge(n1, n2, n3, name, isPassed, address, nil, stu, &stu)
}*/

//------------------------------------------------------
/*type Usb interface {
	Start()
	Stop()
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

type Phone struct {
	name string
}

func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}
func (p Phone) Call() {
	fmt.Println("手机打电话...")
}

type Computer struct{}

func (c Computer) working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	// 遍历多态数组，如果是Phone类型，除了调用接口公用的方法外，再调用它的Call方法
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"尼康"}
	fmt.Println(usbArr)

	var computer Computer
	for _, usb := range usbArr {
		computer.working(usb)
		fmt.Println()
	}
}*/

//-----------------------------------------------
/*func main() {
	var a interface{}
	var num = 8.9
	a = num
	fmt.Println(a)

	// 带检测的类型断言
	if b, ok := a.(float32); ok {
		fmt.Println("convert success,b=", b)
	} else {
		fmt.Println("convert failure")
	}

	fmt.Println("继续执行...")
}*/

/*type Point struct {
	x, y int
}

func main() {
	p := Point{1, 2}
	var a interface{} = p
	fmt.Println(a)
	var b Point
	// b = a
	// 类型断言，如果a是指向Point类型的变量，就转成Point类型，并赋值给b，否则报错
	b = a.(Point)
	fmt.Println(b)
}*/

// 类型断言 语法： b = a.(Point)
// 带检测的类型断言：b, ok := a.(float32)
// 类型判断 switch v.(type)

// 总结：当带有类型断言的方法调用时，指针可以自动解引用，但是对象不能自动转换成指针