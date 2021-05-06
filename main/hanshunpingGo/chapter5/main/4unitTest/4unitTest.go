package main

func AddUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func GetSub(n1, n2 int) int {
	return n1 - n2
}

// 单元测试 框架 testing（针对函数的测试用例和压力测试）
// go test 命令实现单元测试和性能测试

// go test的使用细节
// 1.测试用例文件名以_test.go结尾
// 2.测试用例函数TestXxx() 必须以Test开头，而且Test之后必须是大写字母，不能是小写字母
// 3.TestXxx(t *testing.T) 形参必须是 *testing.T
// 4.一个测试文件中可以有多个测试函数
// 5.go test 命令说明
//     go test -v 不论结果正确或错误，都显示日志
//     go test  只显示错误的结果日志
// 6.输出日志用t.Logf()
// 7.错误日志用t.Fatalf(),并会退出程序
// 8.测试用例函数并没有放在main函数中，能在testing框架中正确执行，这正是测试用例的方便之处
// 9.PASS表示测试用例通过，FAIL表示测试用例未通过
// 10.执行单个测试用例文件，一定要带上被测试的原文件 go test -v 4unitTest_test.go 4unitTest.go
// 11.执行单个测试函数，go test -v -test.run TestAddUpper