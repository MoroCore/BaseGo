package main

import (
	"fmt"
	"math/rand"
	"time"
)

//数组使用的细节
func main() {

	// 1: 数组是多个相同类型数据的集合，一个数组一旦声明，定义，其长度是规定的，不能动态的变化
	var arr01 [3]int //声明并且使用默认初始化,必须指明长度和类型

	arr01[0] = 1
	arr01[1] = 10
	// arr01[2] = 1.1  untyped float constant 类型不匹配
	// arr01[3] = 890 index 3 (constant of type int) is out of bounds 数组下标越界
	fmt.Println(arr01) //[1 10 0]

	//2: var arr []int  这时arr是一个slice切片     不指定长度
	//3: 数组中的元素是可以是任何元素类型，包括值类型和引用类型，但是不能混用
	//4: 数组创建后，如果没有赋值，有默认值  数值 0   字符串 ""   bool false
	//5: 使用数值的步骤: 声明数组并开辟空间  给数组各个元素赋值   使用数组
	//6: 数组的下标从0开始
	var arr04 [3]string
	var index int = 2    //注意  这个编译不报错   运行时  exit status 2
	arr04[index] = "Tom" //因为下标是0-2  所有arr03[3]越界
	fmt.Println("arr04[3]=", arr04[index])
	//7: 数组下标必须在指定范围内使用，否则报 panic：数组越界，比如
	var arr05 [5]int // 则有效下标为 0-4
	fmt.Println(arr05)
	//8: Go 的数组属值类型， 在默认情况下是值传递， 因此会进行值拷贝。数组间不会相互影响
	//9: 如想在其它函数中，去修改原来的数组，可以使用引用传递(指针方式)
	arr06 := [...]int{1, 2, 3}
	test(&arr06)
	fmt.Println("arr06[0]=", arr06[0])

	//10: 长度是数组类型的一部分，在传递函数参数时 需要考虑数组的长度
	charArray()
}
func test(arr *[3]int) {
	(*arr)[0] = 88
}

//1) 创建一个 byte 类型的 26 个元素的数组，分别 放置'A'-'Z‘。使用 for 循环访问所有元素并打印出来。提示：字符数据运算 'A'+1 -> 'B'
func charArray() {

	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i)
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", myChars[i])
	}
}

//2) 请求出一个数组的最大值，并得到对应的下标。
func getMaxIndex() {

	fmt.Println()

	intArr := [...]int{1, -1, 9, 90, 11, 9000}
	maxVal := intArr[0]
	maxValIndex := 0

	for i := 1; i < len(intArr); i++ {

		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("MaxValue=%v MavValIndex=%v ", maxVal, maxValIndex)
}

//3) 请求出一个数组的和和平均值。for-range
func getValue() {

	intArr := [...]int{1, -1, 9, 90, 12}
	sum := 0

	for _, v := range intArr {
		sum += v
	}

	fmt.Printf("sum=%v 平均值=%v ", sum, float64(sum)/float64(len(intArr)))
}

// 4) 要求：随机生成五个数，并将其反转打印
func test1() {

	var intArr [5]int

	len := len(intArr)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		intArr[i] = rand.Intn(100)
	}

	fmt.Println("交换前~===", intArr)

	temp := 0

	for i := 0; i < len/2; i++ {
		temp = intArr[len-1-i]
		intArr[len-1-i] = intArr[i]
		intArr[i] = temp
	}
	fmt.Println("交换后~", intArr)
}
