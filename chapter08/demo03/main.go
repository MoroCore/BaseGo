package main

import "fmt"

func main() {

	//二维数组的定义

	var arr [3][6]int

	fmt.Printf("arr[0]的地址:%p \n", &arr[0])
	fmt.Printf("arr[1]的地址:%p \n", &arr[1])

	fmt.Printf("arr[0][0]的地址:%p \n", &arr[0][0])
	fmt.Printf("arr[1][0]的地址:%p \n", &arr[1][0])

	// 定义二维数组，用于保存三个班，每个班五名同学成绩，并求出每个班级平均分、以及所有班级平均分

	var scores [3][5]float64

	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%d班的第%d个学生的成绩\n", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}
	totalSum := 0.0

	for i := 0; i < len(scores); i++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); i++ {
			sum += scores[i][j]
		}
		totalSum += sum
		fmt.Printf("第%d班级的总分为%v,平均分%v\n", i+1, sum, sum/float64(len(scores[i])))
	}
	fmt.Printf("所有班级的总分为%v,所有班级平均分%v\n", totalSum, totalSum/15)
}
