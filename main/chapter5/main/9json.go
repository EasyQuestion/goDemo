package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// 基本数据类型到 JSON 字符串的编码过程
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	//fmt.Println(slcD)
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	//fmt.Println(mapD)
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// 我们需要提供一个 JSON 包可以存放解码数据的变量。这里的 map[string]interface{} 将保存一个 string 为键，值为任意值的map。
	var dat map[string]interface{}
	// 实际的解码和相关的错误检查。
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	// 将 num 的值转换成 float64类型
	num := dat["num"].(float64)
	fmt.Println(num)

	// 访问嵌套的值需要一系列的转化
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// 我们也可以解码 JSON 值到自定义类型。
	// 这个功能的好处就是可以为我们的程序带来额外的类型安全加强，并且消除在访问数据时的类型断言。
	str := `{"page":1,"fruits":["apple","peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 我们经常使用 byte 和 string 作为使用标准输出时数据和 JSON 表示之间的中间值。
	// 我们也可以和os.Stdout 一样，直接将 JSON 编码直接输出至 os.Writer流中，或者作为 HTTP 响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

// JSON 包可以自动的编码你的自定义类型。编码仅输出可导出的字段，并且默认使用他们的名字作为 JSON 数据的键。
// 你可以给结构字段声明标签来自定义编码的 JSON 数据键名称。
