package main

import (
	"fmt"
	"time"
)
//时间和日期相关函数
func main(){

	//在编程中，程序员会经常使用到日期相关的函数，比如：统计某段代码执行花费的时间等
	//	1)时间和日期相关函数，需要导入time包
	// 	2)time.Time类型，用于表示时间
	// 1: 获取当前时间
	now := time.Now()
	fmt.Printf("now = %v  now type = %T \n", now, now)//now = 2022-05-05 19:29:00.0100446 +0800 CST m=+0.002707701  now type = time.Time

	// 2: 通过now可以获取到年时分秒
	fmt.Printf("年=%v \n",now.Year()) //now.Year() type int 
	fmt.Printf("月=%v \n",now.Month()) //now.Year() type Month 
	fmt.Printf("月=%v \n",int(now.Month())) //now.Year() type int 
	fmt.Printf("日=%v \n",now.Day()) //now.Year() type int 
	fmt.Printf("时=%v \n",now.Hour()) //now.Year() type int 
	fmt.Printf("分=%v \n",now.Minute()) //now.Year() type int 
	fmt.Printf("秒=%v \n",now.Second()) //now.Year() type int 
	
	// 4)格式化日期时间
	//	方式1:就是使用Printf或者SPrintf
	fmt.Printf("当前年月日%d-%d-%d %d:%d:%d \n", now.Year(),now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("当前年月日  %d-%d-%d %d:%d:%d \n", now.Year(),now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Printf("dateStr=%v\n",dateStr)
	//	方式二:使用time.Format()方法完成:
	fmt.Printf(now.Format("2006/01/02 15:01:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006/01/02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()
	//细节 2006/01/02 15:01:05" 规定 否则返回时间不对

	// 5)时间的常量
	//	常量的作用:在程序中可用于获取指定时间单位的时间，比如想得到100毫秒
	// test()

	// 7)time的Unix和UnixNano的方法
	//		Unix将t表示为Unix时间,即从时间点January 1, 1970 UTC到时间点t所经过的时间(单位秒)。
	fmt.Printf("unix时间戳=%v \n",now.Unix())
	//		UnixNano将t表示为Unix时间,即从时间点January 1. 1970 UTC到时间点t所经过的时间(单位纳秒)
	//.     如果纳秒为单位的unix时间超出了int64能表示的范围，结果是未定义的。注意这就意味着Time零值调用UnixNano方法的话,结果是未定义的。
	fmt.Printf("unixNano时间戳=%v",now.UnixNano())   //1s = 10^9ns
}

//需求，每隔1秒中打印一个数字，打印到100时就退出
//需求2:每隔0. 1秒中打印一个数字，打印到100时就退出
func test() {
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
		if i == 100 {
			break
		}
	}
}