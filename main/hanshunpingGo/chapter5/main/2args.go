package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名，默认为空串")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空串")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	// 调用 flag.Parse() 来执行命令行解析
	flag.Parse()

	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v\n",
		user, pwd, host, port)
}

/*func main() {
	fmt.Println("命令行的参数个数为：", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
}*/

// 命令行参数  os.Args
// 命令行参数的解析 flag包
// flag.StringVar()  flag.IntVar()  flag.Parse()
