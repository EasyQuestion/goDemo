package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var ops int64 = 0

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				// 为了确保这个 Go 协程不会在调度中饿死，我们在每次操作后明确的使用 runtime.Gosched()进行释放。
				// 这个释放一般是自动处理的，像例如每个通道操作后或者 time.Sleep 的阻塞调用后相似，
				// 但是在这个例子中我们需要手动的处理。
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}

// 使用原子操作来管理简单的计数器。对于更加复杂的情况，我们可以使用一个互斥锁来在 Go 协程间安全的访问数据。

// 每次循环读取，我们使用一个键来进行访问，
// Lock() 这个 mutex 来确保对 state 的独占访问，读取选定的键的值，
// Unlock() 这个mutex，并且 ops 值加 1。