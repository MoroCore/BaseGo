package main

import "fmt"

// 查找
func main() {

	//在GoLang中 常用两种查找 顺序查找   二分查找

	// 1: 顺序查找
	name := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}

	var heroName = ""

	// fmt.Println("请输入要查找的人名.....")
	fmt.Scan(&name)

	//顺序查找
	for i := 0; i < len(name); i++ {

		if heroName == name[i] {
			fmt.Printf("找到%v ,下标%v \n", heroName, i)
			break
		} else {
			fmt.Printf("没有找到%v \n", heroName)
		}
	}
	//顺序查找 推荐使用
	index := -1
	for i := 0; i < len(name); i++ {

		if heroName == name[i] {
			index = -1
			break
		}

	}

	if index != -1 {
		fmt.Printf("找到%v ,下标%v \n", heroName, index)
	} else {
		fmt.Printf("没有找到%v \n", heroName)
	}

	//对一个有序的数组 进行二分法查找
	arr := [6]int{1, 8, 89, 1000, 1234}
	BinaryFind(&arr, 0, len(arr), 89)

}

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {

	if leftIndex > rightIndex {
		fmt.Printf("找不到")
		return
	}

	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle]  < findVal {
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了,下标为%v \n", middle)
	}
}
