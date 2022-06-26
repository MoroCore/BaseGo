package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

//对文件的整体进行操作
func main() {

	//细节  相对路径 ./是当前项目的目录
	src := "./chapter12/demo04/涂山红红.jpeg"
	CopyFile("./chapter12/demo04/copy涂山红红.jpeg", src)

	charCount := countNumber("./chapter12/demo04/test.txt")
	fmt.Printf("charCount = %v", charCount)
}

//1: 文件拷贝
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {

	file, err := os.Open(srcFileName)

	if err != nil {
		fmt.Printf("open file err = %v", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	openFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("open file  err =  %v", err)
		return
	}
	defer openFile.Close()
	writer := bufio.NewWriter(openFile)

	return io.Copy(writer, reader)
}

// 统计英文 数字 空格 和其他字符数量
type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OhterCount int
}

func countNumber(src string) CharCount {

	file, err := os.Open(src)

	if err != nil {
		fmt.Printf("open file err = %v", err)
		return CharCount{}
	}
	defer file.Close()

	var count CharCount
	reader := bufio.NewReader(file)

	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		runes := []rune(readString)
		for _, v := range runes {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OhterCount++
			}
		}
	}
	return count
}

//3: 命令行参数
func getOrderArgs() {

	//os.Args是一个string切片  用来存储所有的命令行参数
	fmt.Println("命令行参数:", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%v] = %v \n", i, v)
	}
}

// 4: flag包解析命令行参数
//说明:前面的方式是比较原生的方式，对解析参数不是特别的方便，特别是带有指定参数形式的命令行。
//比如: cmd>main.exe -f c:/aaa.txt -p 200 -u root这样的形式命令行，go 设计者给我们提供了flag包，可以方便的解析命令行参数，而且参数顺序可以随意
func userFlagParseArgs() {
	var user string
	var pwd string
	var host string
	var port int
	//&user就是接收用户命令行中输入的  -u  后面的参数值
	//"u”,就是-u指定参数
	// "", 默认值
	//"用户名,默认为空”  说明
	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号,默认为330")
	flag.Parse()

}
