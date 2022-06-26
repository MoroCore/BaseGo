package main

import (
	"fmt"
	"strings"
	"strconv"
)
//字符串常用的系统函数
func main() {

	// 1) 统计字符串的长度，按字节 len(str)
	str := "Hello毕"
	fmt.Println("str len = ",len(str)) //8

	// 2) 字符串遍历，同时处理有中文的问题 r := []rune(str)
	str2:= "Hello毕"
	r := []rune(str2) //int32
	for i := 0 ; i < len(r) ; i++ {
		fmt.Printf("字符=%c \n",r[i])
	}
	// 3) 字符串转整数: n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("转化错误",err)
	}else {
		fmt.Println("转化的结构是",n)
	}

	// 4) 整数转字符串 str = strconv.Itoa(12345)
	n1 := strconv.Itoa(123)
	fmt.Printf("str=%v , str=%T",n1, n1)

	// 5) 字符串 转 []byte: var bytes = []byte("hello go")
	var bytes = []byte("Hello,World")
	fmt.Printf("bytes=%v \n",bytes)

	// 6) []byte 转 字符串: str = string([]byte{97, 98, 99})
	str3 := string([]byte{97,98,99})
	fmt.Printf("str3=%v \n ",str3)

	// 8) 查找子串是否在指定的字符串中: strings.Contains("seafood", "foo") //true
	b := strings.Contains("seeadff","mmm") 
	fmt.Printf("b=%v \n",b)

	// 9) 统计一个字符串有几个指定的子串 ： strings.Count("ceheese", "e") //4
	num := strings.Count("csfsaga","e")
	fmt.Printf("number=%v \n",num)

	// 10) 不区分大小写的字符串比较(==是区分字母大小写的): fmt.Println(strings.EqualFold("abc", "Abc")) // true
	b = strings.EqualFold("abc","abd")
	fmt.Printf("b=%v \n",b)

	fmt.Println("结果","abc"=="abc") //区分大小写

	// 11)返回子串在字符串第一次出现的index值，如果没有返回-1:strings.Index("NLT_abc","abc")//4

	index := strings.Index("aaaaadafaf","aaa")
	fmt.Printf("index=%v \n",index)

	// 12)返回子串在字符串最后一次出现的index，如没有返回-1:strings.LastIndex("gogolang","go")
	lastIndex := strings.LastIndex("aaaagagv","agvaa")
	fmt.Printf("lastIndex=%v \n",lastIndex)

	//13)将指定的子串替换成另外一个子串:strings.Replace("gogohello","go","go语言",n)n可以指定你希望替换几个，如果n=-1表示全部替换
	str4 := "go go hello"
	str = strings.Replace(str4,"go","beijing",-1)

	//14)按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：strings.Split("hello,wrold,ok",",")
	strArr := strings.Split("helagavaabh ahb"," ")
	for i := 0; i < len(strArr) ; i++ {
		fmt.Printf("str[%v]=%v\n",i,strArr[i])
	}

	//15)将字符串的字母进行大小写的转换:strings.ToLower("Go")//gostrings.ToUpper("Go")//GO
	strings.ToLower("GGGGGG")
	strings.ToUpper("aaaaaaaa")

	//16)将字符串左右两边的空格去掉：strings.TrimSpace("tnalonegopherntrn)
	strings.TrimSpace(" agag ")

	//17)将字符串左右两边指定的字符去掉：strings.Trim("!hello!","!")
	strings.Trim("@hellllvga","!")

	//18)将字符串左边指定的字符去掉：strings.TrimLeft("!hello!","!")//["hello"]//将左边!和""去掉
	strings.TrimLeft("ssss","sss")

	//19)将字符串右边指定的字符去掉：strings.TrimRight("!hello!","!")//["hello"]//将右边!和""去掉
	strings.TrimRight("ssss","sss")

	//20)判断字符串是否以指定的字符串开头:strings.HasPrefix("ftp://192.168.10.1","ftp")//true
	strings.HasPrefix("sfaga","Ggg")

	//21)判断字符串是否以指定的字符串结束:strings.HasSuffix("NLT_abc.jpg","abc")//false
	strings.HasSuffix("sfagva","sfagv")
}