package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//写文件
func main() {
	openFileAndReadAllDataAndAppendData()
	//OpenFile函数是一个更广泛使用的函数
	// 参数1: 文件的路径
	// 参数2: 文件打开的方式
	// 	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	// 	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	//	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	//  The remaining values may be or'ed in to control behavior.
	// 	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	// 	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	// 	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	// 	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	// 	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
	// 打开文件的权限
	path := "./test.txt"
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("file open err = %v", err)
		return
	}
	defer file.Close()

	str := "Hello,Carlos \n"

	write := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		write.WriteString(str)
	}
	//因为write是带缓存，因此调用WriteString方法是，其实是先将数据写入缓存区,所以需要调用带Flush方法
	//将缓存的数据真正写入文件，否则文件中没有数据。
	write.Flush()
	openExistFileAndCoverOldData()
}

//2 : 打开一个已存在的文件,用新数据覆盖旧数据
func openExistFileAndCoverOldData() {

	url := "../test1.txt"
	//O_TRUNC 如果可以打开文件 清空文件
	file, err := os.OpenFile(url, os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("file open err = %v", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer.WriteString("moro\n")
	}
	writer.Flush()
}

//3: 打开一个已存在的文件  将原来的内容读出到控制台  并在源文件上追加数据
func openFileAndReadAllDataAndAppendData() {

	path := "./test1.txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(" file open fail , fail reason ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		fmt.Print(readString)
	}
	newData := "Hello,Go Lang \n"
	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer.WriteString(newData)
	}
	writer.Flush()
}
