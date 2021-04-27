package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	// 统计文件中字母、数字、空格、其它字符的个数
	fileName := "d:/test/abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	count := CharCount{}
	for {
		content, err := reader.ReadString('\n')
		for _, v := range content {
			switch {
			case v >= 'a' && v <= 'z', v >= 'A' && v <= 'Z': // 也可以用fallthrough来实现
				count.ChCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			default:
				count.OtherCount++
			}
		}

		if err == io.EOF {
			break
		}
	}

	fmt.Println(count)
}

//------------------------------------------------------------------
/*func CopyFile(dstFilePath, srcFilePath string) (written int64, err error) {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Println("srcFile open err=", err)
		return
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("dstFile open err=", err)
		return
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}

func main() {
	// 二进制文件的拷贝 图片 阿噗壁纸_(1).jpg 拷贝到test文件夹下
	srcFilePath := "d:/阿噗壁纸_(1).jpg"
	dstFilePath := "d:/test/abc.jpg"
	_, err := CopyFile(dstFilePath, srcFilePath)
	if err != nil {
		fmt.Println("copy file err=", err)
	} else {
		fmt.Println("copy successfully...")
	}
}*/

//-----------------------------------------------------
/*func main() {
	// 判断文件或文件夹是否存在
	//fileInfo, err := os.Stat("d:/abc.txt") // 存在的文件
	fileInfo, err := os.Stat("d:/kkk.txt") // 不存在的文件
	if err == nil {
		fmt.Println("file exists,fileInfo=", fileInfo)
	} else if os.IsNotExist(err) {
		fmt.Println("file not exists.err=", err)
	} else {
		fmt.Println("err=", err)
	}
}*/

//----------------------------------------------------------
/*func main() {
	// 一次性将文件复制到另一个文件中
	sourceFilePath := "d:/abc.txt"
	targetFilePath := "d:/test/kkk.txt"

	data, err := ioutil.ReadFile(sourceFilePath)
	if err != nil {
		fmt.Println("read file err=", err)
		return
	}

	err = ioutil.WriteFile(targetFilePath, data, 0666)
	if err != nil {
		fmt.Println("write file err=", err)
	}
	fmt.Println("文件写入完成...")
}*/

//------------------------------------------------------------
/*func main() {
	filePath := "d:/abc.txt"
	// 练习1.创建新文件 d:/abc.txt ,并向文件中写入5次 "hello,Gardon\n"
	//file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)

	// 练习2.已存在文件，覆盖旧内容，写入10次 “你好，尚硅谷”
	// 对于练习2，可以和练习1的写法一样，也可以写成
	// 但是 os.O_TRUNC 会清空文件，所以谨慎使用
	//file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)

	// 练习3.已存在的文件，追加内容"HELLO,ENGLISH"
	//file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)

	// 练习4.已存在的文件，读出文件内容，并追加5句“hello,北京”
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	defer file.Close()

	// 读出文件内容
	reader := bufio.NewReader(file)
	for {
		content, err := reader.ReadString('\n')
		fmt.Print(content)
		if err == io.EOF {
			break
		}
	}
	fmt.Println()
	fmt.Println("文件读取结束...")

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		// 有的编辑器可以不能识别\n,需要写成\r\n才行
		writer.WriteString("hello,北京！\n")
	}
	// bufio.writer 是带有缓存的，写完要把内容从内存刷入到磁盘中
	// 否则文件中没有内容！！！
	writer.Flush()
	fmt.Println("文件写入完成...")
}*/

//--------------------------------------------------------
/*func main() {
	// 读取文件的内容（一次性全部读取到内存中，适合文件较小的情况）ioutil.ReadFile()
	// 没有显式地打开文件，也不需要显式地关闭文件
	filePath := "d:/a.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file err=", err)
	}

	fmt.Println("content:", string(content))
}*/

//------------------------------------------------------
/*func main() {
	// 读取文件的内容（带缓冲区的方式）:适合读取比较大的文件
	file, err := os.Open("d:/a.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer file.Close() // 记得关闭文件

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 遇到换行符就结束，即读取一行内容
		fmt.Print(str)
		if err == io.EOF { // 表示读到文件末尾
			break
		}
	}
	fmt.Println()
	fmt.Println("===================文件读取结束...")
}*/

//--------------------------------------------------
/*func main() {
	// 这里的file ,基本上是以*File 指针的方式来操作文件的
	// file 的叫法：文件对象、文件指针、文件句柄
	// 一定不要忘记关闭文件！！！否则会有内存泄漏问题
	file, err := os.Open("d:/a.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	fmt.Printf("file=%v", file)

	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}*/

// go中文件的使用
// 文件就是数据源
// 文件在程序中以流的形式操作（输入流、输出流）
// os.File（结构体） 封装所有文件相关操作

// 打开文件os.Open()、关闭文件file.Close()

// 读取文件的2种方式
// 读取文件的内容（带缓冲区的方式）os.Open(),file.Close(),bufio.NewReader(),reader.ReadString(),io.EOF
// 读取文件的内容（一次性全部读取到内存中，适合文件较小的情况）ioutil.ReadFile()

// os.OpenFile() 涉及到的参数说明
// 打开文件的几种方式-windows环境下
/*  这几种方式可以通过 | 组合使用
const (
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)
*/
// 打开文件的几种方式-linux/unix环境下:FileMode (可以看linux的权限管理)
// 写入文件的4种方式 os.OpenFile(),file.Close(),bufio.NewWriter(),writer.WriteString(),
//                      os.O_WRONLY | os.O_CREATE
//                      os.O_WRONLY | os.O_TRUNC  这种方式会清空文件，谨慎使用
//                      os.O_WRONLY | os.O_APPEND
//                      os.O_RDWR | os.O_APPEND
// 一次性写入文件 ioutil.WriteFile()

// 文件或目录是否存在 os.Stat()

// 二进制文件的拷贝（图片，音频，视频）：io.Copy()
