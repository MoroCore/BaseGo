package main

import "fmt"

//切片slice
func main() {
	//切片产生的原因:如果数组的长度不确定，如何解决  切片

	//1: 切片的介绍
	//	切片的英文是slice
	//  切片是数组的一个引用，因此切片的类型是引用类型  在参数传递时，遵守传递的机制
	//  切片的使用和数组类似，遍历切片和访问切片 求切片的长度都一样
	//  切片的长度是可变的，因此切片是一个动态的数组
	//	切片的定义   var 切片名 []类型   var a []int
	//2:切面的快速入门
	var intArr = [...]int{1, 22, 33, 66, 99} //int数组
	//声明和定义一个切面
	// 方式1: 在数组的基础上定义切片
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)             //[1 22 33 66 99]
	fmt.Println("slice 的元素是 = ", slice)        //[22 33]
	fmt.Println("slice 的元素个数是 = ", len(slice)) // 2
	fmt.Println("slice 的容量是 = ", cap(slice))   //  4   动态变化
	//3:切面在内存中的存在形式
	//  slice 的确是一个引用类型
	//  slice 从底层来说，其实就是一个数据结构(struct 结构体)
	// type slice struct {
	//	 ptr *[2]int
	// 	 len int
	//   cap int
	// }
	//4: 切片的使用
	//  方案1: 在数组的基础上使用 快速入门案例  缺点: 必须存在一个数组
	//	方案2: make创建切片   var 切片名 []type = make([]type, len, [cap])
	//		解析: type创建什么类型的切片     len必选      cap是可选项  如果选择cap  cap>=len
	var slice1 []float64 = make([]float64, 5)
	fmt.Println("slice1 =", slice1)
	//  通过make方式创建的切片可以指定切片的大小和容量
	//  如果切片没有赋值，那么默认值
	//  make创建切片的底层原理: make创建一个对外不可见的数组  通过slice维护
	//	方案3:  定义一个切片，就直接指定具体数组
	var slice2 []string = []string{"tom"}
	//数组的定义方式
	// var num1 [10]int;
	// var num2 = [3]int;
	// var num3 = [...]int{1}
	// var num4 = [...]int{1:10}
	//     num5 := [3]int;
	fmt.Println("slice2 =", slice2)
	//方式1和方式2的区别  数组是否可见
}
