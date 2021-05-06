package main

import (
	"fmt"
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	if res != 55 {
		t.Fatalf("AddUpper(10) 执行错误，期望结果 %v,实际结果 %v\n", 55, res)
	}
	t.Logf("AddUpper(10) 执行正确...")
}

func TestHello(t *testing.T) {
	fmt.Println("testHello()")
}
