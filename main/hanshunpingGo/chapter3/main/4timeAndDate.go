package main

import (
	"fmt"
	"strconv"
	"time"
)

/*func main() {
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now) // 类型是 time.Time

	fmt.Println("年：", now.Year())
	fmt.Println("月：", int(now.Month()))
	fmt.Println("日：", now.Day())
	fmt.Println("时：", now.Hour())
	fmt.Println("分：", now.Minute())
	fmt.Println("秒：", now.Second())

	// 日期格式化方一：
	fmt.Printf("当前年月日%02d-%02d-%02d %02d:%02d:%02d\n", //换行时，一定要保证结尾有逗号
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
}*/

/*func main() {

	// 每隔0.1秒就打印一次，到100结束
	for i := 0; i < 100; i++ {
		fmt.Println("i=", i, "time=", time.Now())
		time.Sleep(time.Millisecond * 100) //传递的值必须是整数，不能是浮点数
	}
}*/

/*func main() {
	now := time.Now()
	fmt.Println("unix时间戳：", now.Unix(), "unixNano时间戳：", now.UnixNano())
}*/

func test03() {
	for str, i := "", 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	// 统计一个函数执行时长
	startTime := time.Now().Unix()
	test03()
	endTime := time.Now().Unix()
	spendTime := endTime - startTime
	fmt.Println("execute test03() spend time(s):", spendTime)
}

// 日期和时间的常用函数 需要导入time包
// 1.日期时间类型 time.Time
// 2.获取当前时间 time.Now()
// 3.日期时间获取 now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second()
// 4.日期时间格式化 2种方式：
//       方一：直接使用Printf,Sprintf进行格式化输出
//       方二：now.Format("2006/01/02 15:04:05") 这个格式中数字是写死的，否则得不到正确的结果
//             记忆方式：1月2号下午3点4分5秒，06年
// 5.时间单位常量
/*const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)*/
// 6.休眠 time.Sleep(time.Millisecond * 100) //传递的值必须是整数，不能是浮点数
// 7.获取当前unix时间戳和unix纳秒时间戳 Unix()  UnixNano()
//     可以用作获取随机数字：用作random的随机种子
