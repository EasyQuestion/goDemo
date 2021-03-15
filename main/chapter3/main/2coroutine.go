package main

import "fmt"

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	//f("direct")
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	//  Scanln 代码需要我们在程序退出前按下任意键结束。
	fmt.Scanln(&input)
	fmt.Println("done")
}

// Go 协程 在执行上来说是轻量级的线程。
// 使用 go f(s) 在一个 Go 协程中调用这个函数。这个新的 Go 协程将会并行的执行这个函数调用。
// 也可以为匿名函数启动一个 Go 协程。
// 现在这两个 Go 协程在独立的 Go 协程中异步的运行，所以我们需要等它们执行结束。
// 当我们运行这个程序时，将首先看到阻塞式调用的输出，然后是两个 Go 协程的交替输出。
// 这种交替的情况表示 Go 运行时是以异步的方式运行协程的。