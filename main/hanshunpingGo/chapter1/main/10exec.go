package main

import _ "fmt" //可以用_来忽略导入的包

func main() {
	//var n1 int32 = 12
	//var n2 int8
	//var n3 int8
	//n2 = int8(n1) + 127 //编译通过，但会按溢出处理
	//n3 = int8(n1) + 128 //编译不通过，因为128已经超过int8范围

	//fmt.Println("n2=", n2, "n3=", n3)
}

//func main() {
//	var n1 int32 = 12
//	var n2 int64
//	var n3 int8
//	n2 = n1 + 20 //编译不通过
//	n3 = n1 + 20 //编译不通过
//
//	fmt.Println("n2=", n2, "n3=", n3)
//}
