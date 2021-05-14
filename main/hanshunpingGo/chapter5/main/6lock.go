package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	resMap = make(map[int]int64)
	lock   sync.Mutex
)

func factorial(n int) {
	var res int64 = 1
	for i := 1; i <= n; i++ {
		res *= int64(i)
	}
	// 导致的问题1 concurrent map writes
	lock.Lock()
	resMap[n] = res
	lock.Unlock()
}

func main() {
	// 计算1-200的各个数的阶乘，并放入map中，最后显示出来，用goroutine来实现
	// map需要做成全局变量
	for i := 1; i <= 200; i++ {
		go factorial(i)
	}

	// 导致的问题2：手动等待，时间不确定
	time.Sleep(time.Second * 10)
	lock.Lock()
	fmt.Println(resMap)
	lock.Unlock()
}

// 可以通过 go build -race xxx.go来生成.exe文件，执行时会提示出竞争的情况
// 出现的问题 1.多个协程同时向map中写的冲突问题 2.主线程要等待所有协程执行完成，而等待时间无法确定

// 解决思路1：公共变量map加锁：保证同一时间，只有一个协程向map中写 sync包 Mutex 互斥
// 但仍然不完美：主线程不能确定要等待多长时间，所有协程才能执行完成
// 而且有些读的地方也需要加锁，需要仔细分析，增加开发难度
