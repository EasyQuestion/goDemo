package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println("m:",m)

	m["k1"] = 12
	m["k2"] = 3
	fmt.Println("m:",m)
	fmt.Println("len:",len(m))

	c := m["k1"]
	fmt.Println(c)

	delete(m,"k2")

	_,prs := m["k2"]
	fmt.Println("prs:",prs)

	n := map[string]int{"foo":1,"bar":2}
	fmt.Println("n:",n)
}
