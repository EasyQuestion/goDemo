package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A sayOK", a.Name)
}
func (a *A) hello() {
	fmt.Print("A hello")
}

type B struct {
	A
	Name string
}

func (b *B) hello() {
	fmt.Print("B hello")
}

func main() {
	b := &B{}
	//不论匿名结构体字段的首字母是大写还是小写，都能直接访问
	// 如果有相同的字段或方法，编译器采用就近原则
	// 如果想访问匿名结构体的字段或方法，b.A.Name,b.A.SayOk()
	b.Name = "tom"
	b.A.Name = "jack"
	b.age = 12
	//不论匿名结构体方法的首字母是大写还是小写，都能直接访问
	b.SayOk()
	b.hello()
}

//----------------------------------------------------------------
/*type Student struct {
	Name  string
	Age   int
	Score int
}

func (stu *Student) showInfo() {
	fmt.Printf("name=[%v] age=[%v] score=[%v]\n", stu.Name, stu.Age, stu.Score)
}
func (stu *Student) SetScore(score int) {
	stu.Score = score
}

type Pupil struct {
	Student //匿名结构体：因为只有字段类型，没有字段名
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试...")
}

type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试...")
}

func main() {
	// 继承的案例演示
	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 10
	pupil.testing()
	pupil.Student.SetScore(90)
	pupil.Student.showInfo()

	graduate := &Graduate{}
	graduate.Student.Name = "jack"
	graduate.Student.Age = 20
	graduate.testing()
	graduate.Student.SetScore(98)
	graduate.Student.showInfo()
}*/
//--------------------------------------------------
/*import (
	"fmt"
	"goDemo/main/hanshunpingGo/chapter4/model"
)

func main() {
	// 封装的案例演示
	person := model.NewPerson("tom")
	fmt.Println(person)

	person.Name = "jack"
	person.SetAge(20)
	person.SetSalary(5000)
	fmt.Println(person)
	fmt.Println("age=", person.GetAge(), "salary=", person.GetSalary())
}*/

// 面向对象编程思想 ：抽象
// 面向对象编程三大特征：封装、继承、多态
// 1）封装：将抽象出来的字段和对应的操作封装在一起，数据被保护在内部，
//      程序的其它包只能通过被授权的操作（方法）才能对字段进行操作
// 封装的好处：隐藏实现的细节；可以对数据进行安全验证，保证安全合理
// 封装的体现：对结构体中的属性进行封装；通过方法、包，实现封装（struct,首字母大小写+工厂模式）

// 2)继承：当多个结构体有共用的字段和方法时，可以抽象到另一个结构体中
// 继承的好处：代码复用（解决结构体和方法的代码冗余）;提高代码的可维护性和功能的可扩展性
// 继承的体现：匿名结构体
//      （如果一个struct中嵌入了另一个匿名结构体，那么这个struct就能直接访问匿名结构体的字段和方法）

// 使用注意事项：
// 1.结构体可以使用匿名结构体的所有字段和方法（不论首字母是否大小写）
// 2.匿名结构体字段的访问可以简化，当作自己的字段和方法来访问
// 3.如果结构体和匿名结构体有相同的字段或方法，编译器会采用就近原则
// 4.如果有多个匿名结构体，都有相同的字段或方法，需要明确指定是哪个，否则编译器报错
// 5.如果一个结构体内部嵌套了一个有名的结构体，这种模式叫组合，
//                      那么访问组合结构体的字段或方法时，必须带上结构体的名字