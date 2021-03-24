package main

// Go 提供了一个 flag 包，支持基本的命令行标志解析。
import (
	"flag"
	"fmt"
)

func main() {
	// 基本的标记声明仅支持字符串、整数和布尔值选项。
	// 声明一个默认值为 "foo" 的字符串标志 word并带有一个简短的描述。
	// 这里的 flag.String 函数返回一个字符串指针（不是一个字符串值），
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// 用程序中已有的参数来声明一个标志也是可以的。注意在标志声明函数中需要使用该参数的指针。
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	// 所有标志都声明完成以后，调用 flag.Parse() 来执行命令行解析。
	flag.Parse()

	// 这里我们将仅输出解析的选项以及后面的位置参数。
	// 注意，我们需要使用类似 *wordPtr 这样的语法来对指针解引用，从而得到选项的实际值。
	fmt.Println("word:", *wordPtr)
	fmt.Println("num:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
// 命令行标志是命令行程序指定选项的常用方式。例如，在 wc -l 中，这个 -l 就是一个命令行标志。
// 测试这个程序前，最好将这个程序编译成二进制文件，然后再运行这个程序。
// go build command-line-flags.go
// ./command-line-flags -word=opt
// ./command-line-flags -word=opt a1 a2 a3
// ./command-line-flags -word=opt a1 a2 a3 -numb=7
// ./command-line-flags -h  使用 -h 或者 --help 标志来得到自动生成的这个命令行程序的帮助文本。
// ./command-line-flags -wat