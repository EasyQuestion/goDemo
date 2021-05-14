package main

func main() {

}

// channel本质是一个队列 （先进先出）
// channel本身就是线程安全的，多个协程访问时，不需要加锁（编译器在底层维护）
// channel中的数据是有类型限制的