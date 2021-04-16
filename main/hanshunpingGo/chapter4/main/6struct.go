package main

import "fmt"

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

type Person struct {
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
}

// go 通过结构体struct来实现面向对象编程的特性
// golang中 面向接口编程 是非常重要的特性（通过interface接口关联，耦合性低，也非常灵活）
// golang中仍然有继承、封装和多态，只是去掉了extends关键字、方法重载、构造函数、析构函数、隐藏的this指针等

// struct 是值类型，变量直接指向具体的值
// 定义struct的4种方式：
// 1.先声明类型，再每个属性赋值
// 2.monster1 := Monster{Name: "牛魔王", Age: 500}
//   monster1 := Monster{"牛魔王", 500}
//   monster1 := Monster{}
// 3.var person *Person = new(Person) 返回一个指针类型
// 4.var person *Person = &Person{}

// 可以直接使用 person.Name 等价于 (*person).Name go语言会自动解引用
// 但是不能这样写 *person.Name

// struct使用注意事项
// 1.结构体的所有字段在内存中是连续的