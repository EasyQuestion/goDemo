package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"monster_name"`
	Age      int    `json:"monster_age"` // 通过反射机制实现
	Birthday string
	Skill    string
}

// 将结构体序列化成json串
func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "1800-11-11",
		Skill:    "牛魔拳",
	}

	//data, err := json.Marshal(monster)
	// 结构体是值类型，也可以传引用类型，传它的指针
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("marshal json err=", err)
	} else {
		fmt.Println("data=", string(data))
	}
}

func testUnmarshalStruct() {
	str := "{\"monster_name\":\"牛魔王\",\"monster_age\":500,\"Birthday\":\"1800-11-11\",\"Skill\":\"牛魔拳\"}"
	monster := Monster{}
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("unmarshal json err=", err)
	} else {
		fmt.Println("monster=", monster)
	}
}

func testUnmarshalMap() {
	str := "{\"address\":\"洪崖洞\",\"age\":100,\"name\":\"红孩儿\"}"
	var myMap map[string]interface{}
	// 反序列化时，map不需要make空间，因为json.Unmarshal()会给map分配空间
	err := json.Unmarshal([]byte(str), &myMap)
	if err != nil {
		fmt.Println("unmarshal json err=", err)
	} else {
		fmt.Println("myMap=", myMap)
	}
}

func testMap() {
	// 对map进行json序列化
	myMap := make(map[string]interface{})
	myMap["name"] = "红孩儿"
	myMap["age"] = 100
	myMap["address"] = "洪崖洞"
	data, err := json.Marshal(myMap)
	if err != nil {
		fmt.Println("marshal json err=", err)
	} else {
		fmt.Println("data=", string(data))
	}
}

func testSlice() {
	// 对切片Slice进行序列化
	var slice []map[string]interface{}

	map1 := make(map[string]interface{})
	map1["name"] = "jack"
	map1["age"] = 7
	map1["address"] = "北京"
	slice = append(slice, map1)

	map2 := make(map[string]interface{})
	map2["name"] = "tom"
	map2["age"] = 20
	map2["address"] = [2]string{"墨西哥", "夏威夷"}
	slice = append(slice, map2)

	map3 := make(map[string]interface{})
	map3["name"] = "mic"
	map3["age"] = 40
	map3["address"] = "冰岛"
	slice = append(slice, map3)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("marshal json err=", err)
	} else {
		fmt.Println("data=", string(data))
	}
}

func testUnmarshalSlice() {
	str := " [{\"address\":\"北京\",\"age\":7,\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":20,\"name\":\"tom\"}," +
		"{\"address\":\"冰岛\",\"age\":40,\"name\":\"mic\"}]"
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("unmarshal json err=", err)
	} else {
		fmt.Println("slice=", slice)
	}
}

func main() {
	testStruct()
	testMap()
	testSlice()

	testUnmarshalStruct()
	testUnmarshalMap()
	testUnmarshalSlice()
}

// json 序列化，反序列化  json.Mashal()  json.Unmarshal()
// 结构类型:key-value 一般都是对有键值对的数据类型进行序列化，比如 数组，map，切片，结构体struct
// 也可以对基本数据类型进行序列化，但是没有太大意义

// struct结构体内的tag标签，可以在json序列化时指定key的名字
// 通过反射机制来实现

// map在反序列化时，不需要make空间，反序列化方法会分配空间

// 在反序列化时，要确保数据类型一致
