package main

import (
	"encoding/json"
	"fmt"
)

/*type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat2 := Cat{Name: "小花", Age: 100, Color: "花色"}
	cats := make(map[string]Cat)
	cats[cat1.Name] = cat1
	cats[cat2.Name] = cat2
	fmt.Println(cats)

	fmt.Println("请输入小猫的名字：")
	var catName string
	fmt.Scanln(&catName)
	if cat, ok := cats[catName]; ok {
		fmt.Println(cat)
	} else {
		fmt.Println("老太太没有养这只猫...")
	}
}*/
//-------------------------------------
/*type Monster struct {
	Name string
	Age  int
}

func main() {
	monster1 := Monster{Name: "牛魔王", Age: 500}
	monster2 := monster1
	monster2.Name = "牛角大王"
	fmt.Println("monster1=", monster1)
	fmt.Println("monster2=", monster2)
}*/
//---------------------------------------------
/*type Person struct {
	Name string
	Age int
}

func main() {
	p1 := Person{"tom",20}
	p2 := &p1
	p2.Name = "jack"
	fmt.Println("p1=",p1)
	fmt.Println("p2=",p2)
	//fmt.Println(*p2.Name) // 编译出错，不能这样写
}*/
//-----------------------------------------
/*type Point struct {
	x int
	y int
}

type React struct {
	leftUp, rightDown Point
}

type React2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := React{Point{1, 2}, Point{3, 4}}
	// 结构体的所有字段在内存中是连续的
	// 这样取值时会非常快，效率高
	fmt.Printf("r1.leftUp.x的地址%p,r1.leftUp.y的地址%p,r1.rightDown.x的地址%p,r1.rightDown.y的地址%p \n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	r2 := React2{&Point{10, 20}, &Point{30, 40}}
	// 当struct中的内容是指针时，指针所分配的地址仍然是连续的，但是指针的内容有可能不是连续的...
	fmt.Printf("r2.leftUp的地址%p,r1.rightDown的地址%p,\n", &r2.leftUp, &r2.rightDown)
	fmt.Printf("r2.leftUp的值%p,r1.rightDown的值%p,\n", r2.leftUp, r2.rightDown)
}*/
//--------------------------------------------------------------------------
/*type A struct {
	Num int
}

type B struct {
	Num int
	//Color string  不论是A多了Color字段，还是B多个Color字段，AB的强制转换都会编译不通过
}

//type C B // 此时编译器认为B,C是2种不同的数据类型，只能强制类型转换
type C = B // 可以这样，相当于起别名

func main() {
	var a A
	var b B
	a = A(b) // 可以把B类型强制转换成A类型
	fmt.Println(a, b)
	var c C
	c = b
	fmt.Println(c)
}*/
//--------------------------------------------------------------------------
type Monster struct {
	Name   string `json:"name"` // 通过tag标签来实现json的序列化
	Age    int `json:"age"`
	Skills string `json:"skills"`
}

func main() {
	//序列化json和反序列化
	monster := Monster{"铁扇公主", 500, "芭蕉扇"}
	marshal, err := json.Marshal(monster)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("序列化成json的值为%s \n", marshal)
	}
}

// go 通过结构体struct来实现面向对象编程的特性
// golang中 面向接口编程 是非常重要的特性（通过interface接口关联，耦合性低，也非常灵活）
// golang中仍然有继承、封装和多态，只是去掉了extends关键字、方法重载、构造函数、析构函数、隐藏的this指针等

// struct 是值类型，变量直接指向具体的值
// 定义struct的4种方式：
// 1.先声明类型，再每个属性赋值
// 2.monster1 := Monster{  这样就不用和定义字段的顺序保持一致了
//         Name: "牛魔王",
//         Age: 500,
//        }
//   monster1 := Monster{"牛魔王", 500}
//   monster1 := Monster{}
// 3.var person *Person = new(Person) 返回一个指针类型
// 4.var person *Person = &Person{}

// 可以直接使用 person.Name 等价于 (*person).Name go语言会自动解引用
// 但是不能这样写 *person.Name

// struct使用注意事项
// 1.结构体的所有字段在内存中是连续的
// 2.结构体是用户自定义的类型，和其它类型进行转换时需要有 完全相同的字段（名称、个数、类型）
// 3.结构体进行type重新定义后，编译器认为是2个不同的类型，但是可以相互强制转换
// 4.struct每个字段上，可以写上一个tag,该tag可以通过反射获取到，常见的使用场景是序列化和反序列化
