package main

import "fmt"
import "math/rand"
import "time"

// for 循环控制
func main(){

	//Golang 提供 for-range 的方式，可以方便遍历字符串和数组(注: 数组的遍历，我们放到讲数组的时候再讲解) ，案例说明如何遍历字符串

	//变量字符串方式1  传统方式
	var str string = "hello,world!"
	for i := 0 ;  i < len(str) ; i++ {
		//fmt.Printf("%v ",str[i]) //104 101 108 108 111 44 119 111 114 108 100 33 
		//fmt.Printf("%v ",string(str[i])) //h e l l o , w o r l d ! 
		fmt.Printf("%c ",str[i]) //h e l l o , w o r l d ! 
		// 缺点: 字符串含有中文，那么传统的遍历字符串方式，就是错误，会出现乱码
		//原因是传统的对字符串的遍历是按照字节来遍历，而一个汉字在 utf8 编码是对应 3 个字节。
	}
	
	str = str + "背景"
	str2 :=[]rune(str) //
	for i := 0 ; i < len(str2) ; i++ {
		//fmt.Printf("%v ",str2[i]) //104 101 108 108 111 44 119 111 114 108 100 33 32972 26223 
		//fmt.Printf("%v ",string(str2[i])) //h e l l o , w o r l d !背 景 
		//fmt.Printf("%c ",str2[i]) //h e l l o , w o r l d ! 背 景
	}
	//对应for-range遍历方式而言，是按照字符方式遍历。因此如果有字符串有中文，也是ok
	for  i , v := range str {
		fmt . Printf("index=%d, val=%c \n", i, v) //index=10, val=dindex=11, val=! index=12, val=背 index=15, val=景
	}
	test1()
	test2()
	test3()
}

//随机生成1-100的一个数，直到生成了99这个数，看看你一共用了几次?
func test1() {
	//在go中 需要生成一个随机数种子 否则返回的值总是固定的
	//rand的包 math/rand   time  time包 
	//time.Now().Unix() 返回一个从1970 1 1 0:0:0 到现在的秒数
	count := 0
	rand.Seed(time.Now().Unix())
	//fmt.Println("n=",rand.Intn(100)+1) //Intn(100) 0~99
	for {
		n := rand.Intn(100)+1
		fmt.Println("n=",n)
		count++
		if n == 99 {
			break
		}
	}
	fmt.Println("count=",count)
	//细节
	//	(1)break默认会跳出最近的for循环
	//  (2)break后面可以指定标签，跳出标签对应的for循环

	lable:
	for {

		for {

			for {
				fmt.Println("-----------------------")
				break lable
			}
		}
	}
}
//跳转控制语句-continue
func test2() {
	//continue语句用于结束本次循环，继续执行下一次循环。
	//continue语句出现在多层嵌套的循环语句体中时，可以通过标签指明要跳过的是哪一层循环,这个和前面的break标签的使用的规则一样

	//课后练习题(同学们课后自己完成)：
	//某人有100,000元,每经过一次路口，需要交费,
	//规则如下:当现金>50000时,每次交5%
	//当现金<=50000时,每次交1000
	//编程计算该人可以经过多少次路口,使用forbreak方式完成

	var money = 100000.0
	var count = 0 

	for {
		if money > 50000 {
			money = money - (0.05 * money)
			count++
		}
		if money > 0 && money <= 50000 {
			money -= 1000
			count++
		}
		if money <= 0 {
			break 
		}
	}
	fmt.Println("count=",count)
}
//goto基本介绍
func test3(){
	//1)Go语言的goto语句可以无条件地转移到程序中指定的行。
	//2)goto语句通常与条件语句配合使用。可用来实现条件转移，跳出循环体等功能。
	//3)在Go程序设计中一般不主张使用goto语句，以免造成程序流程的混乱，使理解和调试程序都产生困难

	n := 20
	for n < 10 {
		n--
		fmt.Println("ok----")
		if n == 15 {
			goto lable1
		}
	}
	fmt.Println("test")
	lable1:
}