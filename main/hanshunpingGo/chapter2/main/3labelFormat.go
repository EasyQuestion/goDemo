package main

import (
	"fmt"
	//"goDemo/main/hanshunpingGo/chapter2/model"
)

func main() {
	var _ int = 43 // 编译不会报错,但无法使用
	//var a bc = 12
	var num = 10
	var Num = 20
	fmt.Println("num=", num, "Num=", Num)

	// 语法可以通过,但是强烈不推荐这样使用
	var int = 23 // int可以
	fmt.Println("int=", int)
	var float32 = 23.444 // float32也可以
	fmt.Println("float32=", float32)

	//fmt.Println(model.HeroName)
	fmt.Println(heroName) // 同包下的不同文件内容
}

// 标识符:对各种变量\方法\函数等命名时使用的字节序列(字符串)
// 凡是自己能起名字的地方,都叫标识符

// 标识符的命名规则
// 1.由26个英文字母大小写,0-9,_组成
// 2.数字不可以开头
// 3.严格区分大小写
// 4.标识符不能包含空格
// 5.下划线_ 作为占位符使用,称为空标识符,在go中是一个特殊字符(它对应的值会被忽略),不能单独作为标识符使用
// 6.不能以系统保留关键字(一共25个)作为标识符,比如break,if等等

/*25个保留关键字
break	default	func	interface	select
case	defer	go	map	struct
chan	else	goto	package	switch
const	fallthrough	if	range	type
continue	for	import	return	var
*/

// 标识符使用,注意事项
// 1.包名:保持package的名字和目录保持一致,尽量采取有意义的包名,简短,有意义,不要和标准库冲突
// 2.变量名\函数名\常量名:采用驼峰法
// 3.如果变量名\函数名\常量名首字母大写,则可以被其它包访问,如果首字母小写,则只能在本包中使用
// 可以理解成:首字母大写是公有的,首字母小写是私有的

// 疑问:如何引入其它包的变量或方法(变量名或方法名已经首字母大写)
//	-> %GOPATH%/src为默认路径,import 包的相对路径(调用的方式是 包名.标志符或方法 )
// 疑问:如何引入同包内的其它文件的变量或方法
//  -> 同包内的变量或方法不需大写,也不需要引入,在go run xx.go utils2.go 加入用到的文件即可(go build也同理)
// 在新的版本如go 13.0中,会有go mod 专门来设置包的引用问题

/*预定义标识符(36个):包括基础数据类型和系统内嵌函数
append	bool	byte	cap	close	complex	complex64	complex128	uint16
copy	false	float32	float64	imag	int	int8	int16	uint32
int32	int64	iota	len	make	new	nil	panic	uint64
print	println	real	recover	string	true	uint	uint8	uintptr
 */

