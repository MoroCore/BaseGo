package main

import (
	"fmt"
	"io/ioutil"
)

const (
	defaultBufSize = 4096
)

//读文件
func main() {

	// file, err := os.Open("d:/test.txt")

	// if err != nil {
	// 	fmt.Println("open file err = ", err)
	// }

	// defer file.Close()
	// //创建一个 *Reader  是带缓存的  默认4096
	// reader := bufio.NewReader(file)
	// //循环读入文件的内容
	// for {
	// 	//读到 \n 结束
	// 	str, err := reader.ReadString('\n')
	// 	fmt.Println(str)
	// 	if err == io.EOF { //EOF表示文件结尾
	// 		break
	// 	}
	// }
	// fmt.Println("文件读取结束......")

	//使用ioutil.ReadFile一次性将文件读取到位
	file := "d:/test.txt"
	context, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("read file err = ", err)
	}
	//把读取到的内容显示到终端
	//fmt.Printf("%v",content)  []byte
	fmt.Printf("%v \n", context)
	fmt.Printf("%v", string(context))
	//我们没有显式的open文件,因此也不需要显式的close文件
	//因为,文件的open和close被封装到 ReadFile函数内部
}
