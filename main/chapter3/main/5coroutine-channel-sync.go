package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
func main() {
	done := make(chan bool,1)
	go worker(done)

	// 程序将在接收到通道中 worker 发出的通知前一直阻塞。
	// 如果你把 <- done 这行代码从程序中移除，程序甚至会在 worker还没开始运行时就结束了。
	<- done
}
// 我们可以使用通道来同步 Go 协程间的执行状态。这里是一个使用阻塞的接受方式来等待一个 Go 协程的运行结束。
// done 通道将被用于通知其他 Go 协程这个函数已经工作完毕。
