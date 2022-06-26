package main

import (
	"fmt"
	"os"
)

//打开文件关闭文件
func main() {
	//打开文件
	// file 的不同称呼
	//	file   file对象
	//  file file指针 file file文件句柄
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	fmt.Printf("file=%v", file)

	defer file.Close()

	if err != nil {
		fmt.Println("close file err = ", err)
	}
}
