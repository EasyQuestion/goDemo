package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

func test() {
	// defer + recover 来解决协程中的报错，防止所有程序崩溃
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test() err=", err)
		}
	}()
	var myMap map[int]string
	// 此处会报错 panic: assignment to entry in nil map
	myMap[1] = "golang"
}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("main ", i)
	}
}

//------------------------------------------------------
/*func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 3)
	for i := 0; i < 3; i++ {
		stringChan <- "Hello" + fmt.Sprintf("%d", i)
	}

	// 有时管道不能判断什么时候关闭，所以需要解决读数据阻塞的问题
	// select在管道数据读完时，不会报错deadlock，会继续往下判断
	//label1:
	for {
		select {
		case v := <-intChan:
			fmt.Println("receive from intChan value=", v)
		case v := <-stringChan:
			fmt.Println("receive from stringChan value=", v)
		default:
			fmt.Println("no value received,end...")
			// 当用协程时，用return来停止执行
			return
			//break label1 // 不建议用label,
		}
	}
	fmt.Println("end......")
}*/

//------------------------------------------------
/*func main() {
	// chan2只读
	var chan2 <-chan int
	num := <-chan2
	//chan2<-20 不能写
	fmt.Println(num)

	// chan3只写
	var chan3 chan<- int
	chan3 <- 20
	//num2 := <- chan3 不能读
}*/

//-----------------------------------------------------
/*func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	// math.Sqrt() 二次平方根函数
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//func main() {
//	start := time.Now().UnixNano()
//	for i := 0; i < 20000; i++ {
//		if isPrime(i) {
//			fmt.Print(i," ")
//		}
//	}
//	fmt.Println()
//	end := time.Now().UnixNano()
//	fmt.Println("ordinary spend nano time:", end-start)
//}

func main() {
	// 要求统计1-2000000中，素数有哪些（测试数据8000）
	start := time.Now().UnixNano()
	numChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)
	exitChan := make(chan bool, 4)
	go func() {
		for i := 0; i < 20000; i++ {
			numChan <- i + 1
			//fmt.Println("put num:", i+1)
		}
		close(numChan)
	}()

	for i := 0; i < 4; i++ {
		go func() {
			for {
				v, ok := <-numChan
				if !ok {
					break
				}
				if isPrime(v) {
					primeChan <- v
				}
			}
			//fmt.Println("goroutine exit...")
			exitChan <- true
		}()
	}

	//go func() {
	for i := 0; i < 4; i++ {
		<-exitChan
	}
	end := time.Now().UnixNano()
	fmt.Println("goroutine spend nano time:", end-start)
	fmt.Println("primeChan len:", len(primeChan))
	close(primeChan)
	//}()

	//for v := range primeChan {
	//	fmt.Println(v)
	//}
}*/

//------------------------------------------------------------
/*func sum(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	// 一个协程，将1-2000写入numChan管道中
	// 8个协程，从numChan读取数n，计算1+2+...+n的结果，写入resChan管道中
	// 8个协程完成后，遍历resChan,显示结果 res[1]=1...res[10]=55
	numChan := make(chan int)
	resChan := make(chan map[int]int)
	go func() {
		for i := 1; i <= 20; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	for i := 0; i < 8; i++ {
		go func() {
			for {
				v, ok := <-numChan
				if !ok {
					break
				}
				res := sum(v)
				resMap := make(map[int]int, 1)
				resMap[v] = res
				resChan <- resMap
			}
		}()
	}

	var resArr [20]int
	for i := 0; i < 20; i++ {
		v, ok := <-resChan
		if !ok {
			break
		}
		for key, val := range v {
			resArr[key-1] = val
		}
	}
	fmt.Println(resArr)
}*/

//----------------------------------------------------------------------
// 解决之前出现的问题  计算1-200的各个数的阶乘，并放入map中，最后显示出来，用goroutine来实现
/*func main() {
	var resChan = make(chan map[int]int64, 20)
	// 计算1-200的各个数的阶乘，并放入map中，最后显示出来，用goroutine来实现
	for i := 1; i <= 20; i++ {
		go func(n int) {
			var res int64 = 1
			for i := 1; i <= n; i++ {
				res *= int64(i)
			}
			resMap := make(map[int]int64)
			resMap[n] = res
			resChan <- resMap
		}(i)
	}

	var resArr [20]int64
	for i := 0; i < 20; i++ {
		v,ok := <-resChan
		if !ok {
			break
		}
		for key,val := range v {
			resArr[key-1] = val
		}
	}
	close(resChan)
	fmt.Println(resArr)
}*/

//-----------------------------------------------
/*goroutine与channel协作：2个chan管道
一个协程 writeData 向intChan中写值，写入50个值就关闭管道intChan
一个协程 readData  从intChan中读值
再来一个exitChan bool :当readData协程读取完数据后，向exitChan插入一个true，告诉主线程可以结束了
*/
/*func main() {
	intChan := make(chan int)
	exitChan := make(chan bool)
	go func() {
		for i := 0; i < 50; i++ {
			intChan <- i
			fmt.Println("write data :", i)
		}
		close(intChan)
	}()

	go func() {
		for v := range intChan {
			fmt.Println("read data :", v)
		}
		exitChan <- true
		close(exitChan)
	}()
	<-exitChan
	fmt.Println("ok...")
}*/

//-------------------------------------------------

/*type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	personChan := make(chan Person, 10)
	for i := 0; i < 10; i++ {
		personChan <- Person{fmt.Sprint("tom%d", i), rand.Intn(100) + 1, fmt.Sprintf("laiyinlu%d", i)}
	}
	// 需要先关闭管道，再遍历读值
	close(personChan)
	// 遍历方式1：for+break读值
	//for {
	//	v, ok := <-personChan
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(v)
	//}

	// 遍历方式2：for-range
	for v := range personChan {
		fmt.Println(v)
	}
}*/

//--------------------------------------------------
/*type Cat struct {
	Name  string
	Hobby string
}

func main() {
	myChan := make(chan interface{}, 3)
	myChan <- Cat{Name: "小花", Hobby: "捉小虫"}
	myChan <- Cat{Name: "大白", Hobby: "睡觉"}
	myChan <- 10
	cat1 := <-myChan
	fmt.Printf("cat1= %v(%T)\n", cat1, cat1) // cat1= {小花 捉小虫}(main.Cat)
	//fmt.Println("cat1.Name=", cat1.Name) // 编译不通过
	fmt.Println("cat1.Name=", cat1.(Cat).Name) // 需要类型断言下
}*/

//-------------------------------------------------------------
/*func main() {
	intChan := make(chan int, 3)
	// chan其实是引用类型，它的值存的是内存地址
	fmt.Printf("intChan的值=%v,intChan本身的地址=%p\n", intChan, &intChan)

	intChan <- 10
	num := 211
	intChan <- num
	// 管道的长度，容量
	// 向管道中填值，不能超过容量，否则会报错 all goroutines are asleep - deadlock!
	//intChan<-22
	//intChan<-99
	fmt.Printf("intChan len:%v cap:%v\n", len(intChan), cap(intChan))

	num2 := <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("intChan len:%v cap:%v\n", len(intChan), cap(intChan))
}*/

// channel本质是一个队列 （先进先出）
// channel本身就是线程安全的，多个协程访问时，不需要加锁（编译器在底层维护）
// channel中的数据是有类型限制的（基本数据类型、map、结构体、指针、interface{}（空接口）任意类型等等）

// chan 关键字
// channel是引用类型，必须初始化后才能写入数据，即make后才能使用
// 向管道中写值，不能超过容量，否则会报错
// 管道的价值在于一边存数据，一边取数据
// 在没有使用协程的情况下，如果管道中的值被读完，再读时会报错

// chan 的关闭 close()内置函数
//            关闭后，不能再向管道中写值，但可以从管道中读值
// chan 的遍历 需要先关闭管道再遍历，否则会报错deadlock
//            2种遍历方式：1) for + break 2) for-range
//            不能直接用i< len(intChan)来判断for循环，因为intChan的长度可能会有变化,而且只能取出一半的值

/*goroutine与channel协作：基本思路：2个chan管道
一个协程 writeData 向intChan中写值，写入50个值就关闭管道intChan
一个协程 readData  从intChan中读值
再来一个exitChan bool :当readData协程读取完数据后，向exitChan插入一个true，告诉主线程可以结束了
*/

//  chan 的阻塞：如果管道中，只有读的协程、写的协程中的一个，编译器会判断出deadlock
//                       只要同时有读的协程，写的协程，只会阻塞，但是程序不会出现deadlock

// channel的使用注意事项
// 1.channel可以声明为只读或只写 var chan2 chan<- int 只写不能读 ；var chan2 <-chan int 只读不能写
//         也可以在方法或函数的入参处声明使用
// 2.select  可以解决从管道取数据阻塞的问题
//         基本语法  类似switch
// 3.defer + recover 进行协程中panic的捕获
