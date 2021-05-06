package main

import (
	"testing"
)

func TestGetSub(t *testing.T) {
	res := GetSub(10, 3)
	if res != 7 {
		t.Fatalf("GetSub(10,3) 执行错误，期望结果 %v,实际结果 %v\n", 7, res)
	}
	t.Logf("GetSub(10,3) 执行正确...")
}
