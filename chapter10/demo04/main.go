package main

import "fmt"

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func main() {

	circle := Circle{10}

	area := circle.area()

	fmt.Println("area:", area)

	methodUtils := MethodUtils{}
	methodUtils.Print()
	methodUtils.Print2(3, 10)

	area1 := methodUtils.area(10, 10)
	fmt.Println("area1:", area1)

}

//1)编写结构体(MethodUtils)，编程一个方法，方法不需要参数，在方法中打印一个10*8的矩形，在main方法中调用该方法。
type MethodUtils struct {
}

func (mu MethodUtils) Print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//2)编写一个方法，提供m和n两个参数，方法中打印一个m*n的矩形
func (mu MethodUtils) Print2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//3)编写一个方法算该矩形的面积(可以接收长len，和宽 width)，将其作为方法返回值。在main方法中调用该方法，接收返回的面积值并打印。
func (mu MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

//判断一个数是奇数还是偶数
func (mu MethodUtils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Println(num, "是偶数....")
	} else {
		fmt.Println(num, "是数....")
	}
}
func (mu MethodUtils) zhangzhi(nums [3][3]int) {

	temp := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			temp = nums[j][i]
			nums[j][i] = nums[i][j]
			nums[i][j] = temp
		}
	}
}
