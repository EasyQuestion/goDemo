package main

import "fmt"

var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

func main(){
	for i:=0;i< len(balance);i++{
		fmt.Println(balance[i]," ")
	}
}

//行数
/*const LINES int = 10

// 杨辉三角
func ShowYangHuiTriangle() {
	nums := []int{}
	for i := 0; i < LINES; i++ {
		//补空白
		for j := 0; j < (LINES - i); j++ {
			fmt.Print(" ")
		}
		for j := 0; j < (i + 1); j++ {
			var length = len(nums)
			var value int
			if j == 0 || j == i {
				value = 1
			} else {
				value = nums[length-i] + nums[length-i-1]
			}
			nums = append(nums, value)
			fmt.Print(value, " ")
		}
		fmt.Println("")
	}
}

func main() {
	ShowYangHuiTriangle()
}*/

/*func main() {
	var arr [10]int
	arr[0] = 1
	arr[1] = 1
	arr[2] = 1
	arr[3] = 1
	fmt.Printf("%d,%d,%d,%d", arr[0], arr[1], arr[2], arr[3])
}*/

/*func main() {
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Print(i, " ")
		}
	}
}*/

/*func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf(" %d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
}*/

/*func main() {
	var C, c int
	C = 1
A:
	for C < 100 {
		C++
		for c = 2; c < C; c++ {
			if C%c == 0 {
				goto A //因子不是素数
			}
		}
		fmt.Print(C, " ")
	}
}*/

/*func main(){
	for true{
		fmt.Println("这是无限循环。\n")
	}
}*/

/*func main() {
	const (
		a = iota //0
		b       //1
		c       //2
		d = "ha" //独立值，iota += 1
		e       //"ha"   iota += 1
		f = 100  //iota +=1
		g       //100  iota +=1
		h = iota //7,恢复计数
		i       //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}*/

/*const (
	i=1<<iota
	j=3<<iota
	k
	l
)
const a = iota
const b = iota

func main() {
	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
	fmt.Println("a=",a)
	fmt.Println("b=",b)
}*/
