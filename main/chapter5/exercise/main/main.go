package main

import "fmt"

func init() {
	fmt.Println("init2...")
}

func init() {
	fmt.Println("init1...")
}

func main() {
	flag := true
	if flag == false {
		fmt.Println("flag assert")
	}
}