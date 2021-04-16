package main

import (
	"fmt"
	"sort"
)

/*func main() {
	a := make(map[string]string)
	fmt.Println(a)
	a["宋江"] = "水浒传"
	a["林黛玉"] = "红楼梦"
	fmt.Println(a)
	delete(a, "宋江")
	fmt.Println(a)

	val, isHave := a["宋江"]
	if isHave {
		fmt.Println("val=", val)
	} else {
		fmt.Println("没有这个数据")
	}
}*/

/*func main() {
	// 用map来保存学生的姓名和性别
	stuMap := make(map[string]map[string]string)
	stuMap["stu01"] = make(map[string]string, 2)
	stuMap["stu01"]["name"] = "张三"
	stuMap["stu01"]["sex"] = "男"
	stuMap["stu02"] = make(map[string]string, 2)
	stuMap["stu02"]["name"] = "李四"
	stuMap["stu02"]["sex"] = "女"
	fmt.Println(stuMap)
	fmt.Println(stuMap["stu02"])
}*/
/*func main() {
	cities := map[string]string{"no1": "北京", "no2": "天津", "no3": "上海"}
	for k, v := range cities {
		fmt.Printf("k=%v v=%v ", k, v)
	}
	fmt.Println("cities len:", len(cities))
}*/
//------------------------------------------------
/*func main() {
	monsters := make([]map[string]string, 2)

	monsters[0] = make(map[string]string, 2)
	monsters[0]["name"] = "牛魔王"
	monsters[0]["age"] = "500"

	monsters[1] = make(map[string]string, 2)
	monsters[1]["name"] = "玉兔精"
	monsters[1]["age"] = "400"

	monster := make(map[string]string, 2)
	monster["name"] = "狐狸精"
	monster["age"] = "300"
	// 注意：切片只能用append()来追加，下标赋值时不能超过定义的初始长度，否则会报index out of range
	monsters = append(monsters, monster)

	fmt.Println(monsters)
	fmt.Println("len:", len(monsters))
}*/
//----------------------------------------------------
func main() {
	map1 := make(map[int]int)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	fmt.Println(map1)

	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Print(map1[v], " ")
	}
}

// map :key-value的数据结构 又叫字段或关联数组
// 基本语法 var myMap = make(map[keyType]valueType)
// keyType:valueType:可以是数字类型、string,bool,chann,指针,接口、结构体、数组，只要是能==比较的，都可以
// 声明var myMap = map[keyType]valueType 并不会分配内存，需要make分配内存,之后才能赋值和使用
// map的3种使用方式：
//   方一：先声明再赋值
//   方二：var myMap = make(map[string]string)
//   方三：var myMap = map[string]string{"no1":"成都"}
// delete()内置函数进行map的删除
// 没有一次全部删除的方法，需要遍历删除，或者重新赋值（让GC回收旧数据）
// map中是否存在某个key
// map遍历:只能用for-range
// len()内置函数：统计map的长度

// map切片：[]map[string]string
// 声明语句： make([]map[string]string)

// map排序：map默认key是无序的（不会按输入顺序，也不排序）
// 排序思路：将key存储到一个切片中，对切片进行排序，再按顺序输出对应的value

// map使用注意事项：
// 1.map是引用类型，遵循引用传递机制，会直接修改元素的值
// 2.map容量满了之后会自动扩容，不会panic,能动态增长
// 3.map的value也经常使用struct类型，比value是map类型更好管理