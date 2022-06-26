package main

import "fmt"

func main() {
	//切片的遍历
	// for循环常规方式遍历
	// for-range结构遍历切片
	slice := make([]int, 5)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}
	fmt.Println("----------------------------------")
	for i, v := range slice {
		fmt.Printf("slice[%v]=%v \n", i, v)
	}
	//切片的使用的注意事项和细节讨论
	// 1)切片初始化时varslice=arr[startIndex:endIndex]
	// 2)切片初始化时，仍然不能越界。范围在[0-len(arr)]之间，但是可以动态增长.
	// 3)cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素。
	// 4)切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者make一个空间供切片来使用
	// 5)切片可以继续切片
	slice1 := slice[0:1]
	slice[0] = 100
	fmt.Println("slice[0]=", slice[0])
	fmt.Println("slice1[0]=", slice1[0])
	// 6)用append内置函数，可以对切片进行动态追加
	slice[1] = 200
	slice[2] = 300
	slice = append(slice, 100, 200, 300) //在slice后追加  小技巧  在len后追加
	fmt.Println("slice=", slice)
	slice = append(slice, slice...) //切片后添加切片
	fmt.Println("slice=", slice)
	//	append原理: 对数组的扩容
	// 7)切片的拷贝操作
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println("slice04=", slice4) //[1 2 3 4 5]
	fmt.Println("slice05=", slice5) //[1 2 3 4 5 0 0 0 0 0]
	//分析 copy(para1,para2)参数的数据类型是切片
	//	   按照上面的代码来看,slice4和slice5的数据空间是独立，相互不影响，也就是说slice4[0]=999,slice5[0]仍然是1
	//8 tring和slice
	//1)string底层是一个byte数组，因此string也可以进行切片处理
	str := "hello@Hesss"
	slice6 := str[6:]
	fmt.Println("slice6 =", slice6)
	// 3)string是不可变的，也就说不能通过str[0]='z'方式来修改字符串
	// 4)如果需要修改字符串，可以先将string->[]byte/或者[]rune->修改->重写转成string  rune = int32
	arr01 := []byte(str)
	arr01[0] = 'z'
	str = string(arr01)

	arr02 := []rune(str)
	arr02[0] = 'z'
	str = string(arr02)
}
