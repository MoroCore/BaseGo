package main

import "fmt"

//数组可以存放多个同一类型数据。数组也是一种数据类型，在 Go 中，数组是值类型。
func main() {

	var hens [4]float64

	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4

	totalWeght := 0.0

	for i := 0; i < len(hens); i++ {
		totalWeght += hens[i]
	}

	avgWeight := fmt.Sprintf("%.2f", totalWeght/float64(len(hens)))

	fmt.Printf("totalWeght=%v avgWeight=%v", totalWeght, avgWeight)

	//数组的定义和内存布局
	//var 数组名 [数组大小]数据类型
	// var a [5]int

	//1)数组的地址可以通过数组名来获取&intArr
	//2)数组的第一个元素的地址，就是数组的首地址
	//3)数组的各个元素的地址间隔是依据数组的类型决定，比如int64->8 int32->4

	var intArr [3]int
	fmt.Println(intArr)
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30

	fmt.Println(intArr)
	fmt.Printf("inArr的地址=%p intArr[0] 地址%p  intArr[1] 地址%p intArr[2] 地址%p", &intArr, &intArr[0], &intArr[1], &intArr[2])

	//数组的使用 数组名[下标]
	// var score [5]float64

	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("请输入第%d个元素的值\n", i+1)
	// 	fmt.Scanln(&score[i])
	// }

	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("score[%d]=%v \n", i, score[i])
	// }

	//四种初始化数组的方式
	// 声明时   指定数组类型和空间大小
	var number01 [3]int = [3]int{1, 2, 3}
	fmt.Println("number01 = ", number01)
	//  声明变量   在赋值时指定数组类型和空间大小
	var number02 = [3]int{5, 6, 7}
	fmt.Println("number02 = ", number02)
	//     		 根据赋值的大小  推导空间大小 [.......]
	var number03 = [...]int{8, 9, 10}
	fmt.Println("number03 = ", number03)
	//		:=   代替声明
	number04 := [...]string{"1", "2", "3", "毕亚军"}
	fmt.Println("number04 = ", number04)

	//特殊的赋值方式  指定元素的下标
	number05 := [...]int{1: 10, 2: 20}
	fmt.Println("number05 = ", number05)

	//数组的遍历
	//方案1    for 循环变量  缺点：变量汉字出错
	//	原因：按字节变量  1个字母1个字节  一个汉字3个字节
	for i := 0; i < len(number03); i++ {
		fmt.Printf("number03[%d],值=%d \n", i, number03[i])
	}
	//对for-range 的循环的说明
	//	第一个返回值是 数组的下标
	//  第二个返回值是 数组的元素
	//	i v 变量都是for循环内部可见
	//	如果for循环内部不使用 返回值 可设为_
	//  返回值变量名可指定
	for i, v := range number04 {
		fmt.Printf("number03[%d],值=%s \n", i, v)
	}

}
