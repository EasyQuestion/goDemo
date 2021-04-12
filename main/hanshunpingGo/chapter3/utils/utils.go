//package abc // 一般是小写
package utils
import "fmt"

var Name string
var Age = test2()

func init() {
	Name = "tom~"
	Age = 20
	fmt.Println("utils.init()...")
}

func test2() int {
	fmt.Println("utils.test2()...")
	return 12
}

func Cal(n1 float64, n2 float64) float64 { // 只有首字母大小，函数才是可导出的，才能被其它包引用
	return n1 + n2
}
