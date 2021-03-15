package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

// 定时器 是当你想要在未来某一刻执行一次时使用的
// 打点器 则是当你想要在固定的时间间隔重复执行准备的。
// 打点器将定时的执行，直到我们将它停止。

// 打点器和定时器的机制有点相似：一个通道用来发送数据。
// 打点器可以和定时器一样被停止。一旦一个打点停止了，将不能再从它的通道中接收到值。