package main // 打包
import "fmt" // 引入包

func main() {

	fmt.Println("hello") // 通过“包名.函数名”来调用包内的函数

}
/* 函数的基本语法
func 函数名（形参列表）（返回值列表）{
	执行语句
	return 返回值列表
}
// 返回值列表可以没有
 */

// 包：将变量和函数分散到不同文件中去，也可以解决变量和函数重名的问题，还可以控制变量和函数的访问范围
// （本质就是文件夹）
// 函数名大写，表示其它包能访问，也叫“该函数可导出”

// 函数 vs 方法
