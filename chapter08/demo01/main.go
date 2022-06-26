package main

import "fmt"

//排序  排序是查找的基础
func main() {

	//排序是将一组数据,安装指定的顺序进行排序的过程
	//	排序的分类:
	//		内部排序: 指将需要处理的数据都加载的内部存储器中进行排序
	//			交换式排序发  选择式排序法   插入式排序法
	//		外部排序: 数据量过大,无法将数据全部加载的内存中 需要借助外部存储进行排序
	//			合并排序法     直接合并排序法
	// 冒泡排序是典型的交换式排序
	arr := [5]int{24, 69, 88, 57, 13}
	BubbleSort(&arr)

	fmt.Println("main arr = ", arr)
}
func BubbleSort(arr *[5]int) {

	fmt.Println("arr排序之前=", (*arr))

	temp := 0

	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}

	fmt.Println("arr排序后=", (*arr))
}
