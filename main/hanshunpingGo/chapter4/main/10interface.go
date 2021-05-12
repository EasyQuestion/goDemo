package main

import "fmt"

type TransInfo struct {}
type Fragment interface {
	Exec(transInfo *TransInfo) error
}
type GetPodAction struct {
}
func (g GetPodAction) Exec(transInfo *TransInfo) error {
	//...
	return nil
}

func main() {
	var fragment1 Fragment = new(GetPodAction)
	fmt.Println(fragment1) // &{}

	var fragment2 Fragment = &GetPodAction{}
	fmt.Println(fragment2) // &{}

	var fragment3 Fragment = GetPodAction{}
	fmt.Println(fragment3) // {}

}
//-------------------------------------------------
/*
type Mon struct {
	Name string
	Age  int
}

type MonSlice []Mon

func (ms MonSlice) Len() int {
	return len(ms)
}
func (ms MonSlice) Less(i, j int) bool {
	return ms[i].Age <= ms[j].Age
}
func (ms MonSlice) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func main() {
	ms := MonSlice{}
	for i := 0; i < 10; i++ {
		mon := Mon{}
		mon.Name = fmt.Sprintf("英雄%d", i)
		mon.Age = rand.Intn(100)
		ms = append(ms, mon)
	}

	for _, v := range ms {
		fmt.Println(v)
	}

	fmt.Println("排序后--------------------")
	sort.Sort(ms)
	for _, v := range ms {
		fmt.Println(v)
	}
}*/
//---------------------------------------------------
/*type Usb interface {
	Say()
}

type Stu struct{}

func (stu *Stu) Say() {
	fmt.Println("Say()")
}

func main() {
	var stu Stu = Stu{}
	// 这里是指针类型*Stu实现了接口Usb，所以必须对应好类型
	//var u Usb = stu
	var u Usb = &stu
	fmt.Println("here", u)
}*/
//-----------------------------------------
/*type AInterface interface {
	test01()
	test02()
}

type BInterface interface {
	test01()
	test03()
}

// 一个接口同时继承了2个接口，但是2个接口中有重名的方法，那么编译器编译不通过
type CInterface interface {
	AInterface
	BInterface
}*/

//-----------------------------------
/*type AInterface interface {
	test01()
	test02()
}

type BInterface interface {
	test01()
	test03()
}

// Model同时实现了2个接口，这2个接口中有重名的方法，是可行的
type Model struct{}

func (m Model) test01() {}
func (m Model) test02() {}
func (m Model) test03() {}

func main() {
	m := Model{}
	var a AInterface = m
	var b BInterface = m
	fmt.Println("ok~", a, b)
}*/

//----------------------------------------------------
/*type T interface{} // 空接口

type Model struct{}

func main() {
	var t T
	t = Model{}
	fmt.Println(t)

	var t2 interface{} // 空接口也可以这样定义
	t2 = Model{}
	fmt.Println(t2)

	// 我们可以把任何变量赋值给空接口类型
	var num float64 = 8.8
	t2 = num
	t = num
	fmt.Println(t, t2)
}*/

//--------------------------------------------------
/*type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}

// 如果A接口同时继承了B、C接口，那么必须实现A、B、C接口的所有方法，才算实现了A接口
type AInterface interface {
	BInterface
	CInterface
	test03()
}

type Model struct{}

func (p Model) test01() {}
func (p Model) test02() {}
func (p Model) test03() {}

func main() {
	p := Model{}
	var a AInterface = p
	a.test01()
	a.test02()
	a.test03()
}*/

//---------------------------------------
/*type AInterface interface {
	Say()
}
type BInterface interface {
	Hello()
}

// 一个自定义类型，可以实现多个接口
type Stu struct{}

func (stu Stu) Say() {
	fmt.Println("Stu say()...")
}
func (stu Stu) Hello() {
	fmt.Println("Stu hello()...")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer say()", i)
}

func main() {
	var stu = Stu{}
	var a AInterface = stu
	a.Say()
	var b2 BInterface = stu
	b2.Hello()
	// 只要是自定义类型，都可以实现接口，不仅仅是struct结构体
	var b integer
	b.Say()
}*/

//------------------------------------------------------
// interface中声明\定义了一组方法，并且不需要实现方法的具体内容
// 接口中不能包含任何变量
// 只要实现了接口的所有方法，即为实现了这个接口，不需要显式地实现这个接口
// 因此在golang中，没有implements关键字
// 这种方式，可以实现多个接口，而java中无法做到
/*type Usb interface {
	Start()
	Stop()
}

type Usb2 interface {
	Start()
	Stop()
}

type Camera struct{}

func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

type Phone struct{}

func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

type Computer struct{}

func (c Computer) working(usb Usb2) {
	usb.Start()
	fmt.Println("-------------------------")
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	// 实现了接口Usb的所有方法，golang就认为它实现了Usb接口
	// golang中通过这种方式实现 多态
	// 接口体现了编程思想：高内聚低耦合
	computer.working(phone)
	camera := Camera{}
	computer.working(camera)
}*/

// interface
// 实现了接口Usb的所有方法，golang就认为它实现了Usb接口
// golang中通过这种方式实现 多态
// 体现了编程思想：高内聚低耦合
// 这种方式，可以实现多个接口，而java中无法做到

// interface使用的注意事项
// 1.接口本身不能创建实例，只能指向一个实现了该接口的其他类型的变量（里氏代换原则）
// 2.接口内的方法都没有方法体，只有方法定义
// 3.一个类型实现了一个接口的所有方法，即这个类型实现了这个接口
// 4.只要是自定义类型，都可以实现接口，不仅仅是struct结构体
// 5.一个自定义类型，可以实现多个接口
// 6.golang接口中不能有任何变量
// 7.如果A接口继承了B、C接口，那么必须同时实现A、B、C接口的所有方法，才能算实现了A接口
// 8.interface类型默认是一个指针（引用类型），默认值为nil
// 9.空接口interface{}没有任何方法，所有类型都实现了空接口，我们也可以把任何变量赋值给空接口类型

// 如果是实现多个接口，那么多个接口中可以存在重名的方法
// 如果是接口继承多个接口，那么同时继承的多个接口中，不能存在重名的方法

// 接口 vs 继承
// 接口可以看作是对继承的补充：不破坏继承的结构，同时对方法有一个规范的作用
// 继承解决了代码的复用性和可维护性问题，而接口，是设计上的一种规范，解决了解耦的问题