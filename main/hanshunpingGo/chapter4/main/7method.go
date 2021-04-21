package main

import "fmt"

/*type Person struct {
	Name string
}

//func (p Person) test() {
func (p *Person) test() {
	p.Name = "jack"
	fmt.Println("test() ", p.Name)
}

func main() {
	p := Person{}
	p.Name = "tom"
	p.test()
	fmt.Println("main() ", p.Name)
}*/
/*type integer int
func (i integer) print() {
	fmt.Println(i)
}

func main() {
	var i integer = 10
	i.print()
}*/
//-------------------------------------------------
/*type Student struct {
	Name string
	Age  int
}

func (stu *Student) String() string {
	stu.Name = "tom"
	fmt.Println("get String()...")
	return fmt.Sprintf("Name=%v Age=%v", stu.Name, stu.Age)
}

func main() {
	stu := Student{"jack", 12}
	fmt.Println(&stu)
}*/
//------------------------------------------------
/*type MethodUtils struct {
	// 结构体内可以没有任何字段
}

func (md MethodUtils) Print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func (md MethodUtils) Print2(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func (md MethodUtils) area(length, width float64) float64 {
	return length * width
}

// 编写一个结构体MethodUtils ,给它编写一个方法,返回10*8的矩形,并在main函数中调用
func main() {
	mu := MethodUtils{}
	mu.Print()
	fmt.Println()
	mu.Print2(2, 6)
	fmt.Println()
	fmt.Println("area=", mu.area(2.5, 2))
}*/
//----------------------------------------------------
/*func main() {
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(data)
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data[i]); j++ {
			data[i][j], data[j][i] = data[j][i], data[i][j]
		}
	}
	fmt.Println()
	fmt.Println(data)
}*/
//----------------------------------------------------
type Person struct {
	Name string
	Age int
}
func (p Person) Print() {
	p.Name = "jack"
	fmt.Println("print() ",p.Name)
}
func (p *Person) Print2() {
	p.Name = "mary"
	fmt.Println("print() ",p.Name)
}

func main() {
	// 方法比函数多了一个变量类型,这个变量类型有优化,指针和变量可以用简写的方式来调用
	p := Person{"tom",12}
	p.Print()
	fmt.Println("main() ",p.Name)
	(&p).Print()
	fmt.Println("main() ",p.Name)

	p.Print2()
	fmt.Println("main() ",p.Name)
	(&p).Print2()
	fmt.Println("main() ",p.Name)
}

// go中的方法：作用在指定的数据类型上(是和数据类型绑定的)
// 自定义的类型，都可以有方法，不仅仅是struct

// 方法的调用和传参机制和函数基本一样，唯一多的是绑定了数据类型，会把调用方法的变量也当作实参传递给方法
// 如果调用方法的变量是值类型，会进行值拷贝，如果是引用类型，则会进行引用传递
// 所以在方法上的常用情况 func (p *Person) test(){...

// 方法使用的注意事项
// 1.结构体是值类型,在方法调用时,会进行值拷贝
// 2.如果希望在方法体修改结构体的值,可以通过结构体指针的方式来处理
// 3.golang中的方法作用在具体的数据类型上,和指定的数据类型绑定,
//            因此自定义类型,都可以有方法,不仅仅是struct,还有int,float32等待
// 4.方法的访问范围控制规则,和函数一样,首字母大写,可以在本包和其他包访问;首字母小写,只能在本包内访问
// 5.如果一个类型实现了String()方法,那么fmt.Println()默认会调用变量的String()方法进行输出
//
// 方法与函数的区别
// 1.调用方式不同:方法:数据类型对应的变量.方法(),函数:包名.函数名()
// 2.方法比函数多了一个变量类型,这个变量类型有优化,指针和变量可以简写的方式来调用
