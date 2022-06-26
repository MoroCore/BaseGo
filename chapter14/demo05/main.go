package main

import "fmt"

/*
 * @GoName demo05
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022/6/26 10:11
 */
func main() {
	//一般存放性数据

	//1：存放int
	var intChan chan int
	intChan = make(chan int, 3)
	intChan <- 10
	intChan <- 20
	intChan <- 30
	//因为 intChan的容量是 3  intChan <- 40 deallock
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	//因为 intChan 已经没有数据  num4 := <-intChan
	fmt.Printf("num1=%v num2=%v num3=%v \n", num1, num2, num3)

	//2: 创建一个mapChan 最多可以存放 10个map
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)

	m1 := make(map[string]string)
	m1["city1"] = "北京"
	m1["city2"] = "天津"

	m2 := make(map[string]string)
	m2["hero1"] = "宋江"
	m2["hero2"] = "无伤"
	mapChan <- m1
	mapChan <- m2

	//3: 存放结构体
	var carChan chan Cat
	carChan = make(chan Cat, 10)
	cat1 := Cat{Name: "tom", Age: 15}
	cat2 := Cat{
		Name: "carlos",
		Age:  20,
	}
	carChan <- cat1
	carChan <- cat2

	//取出结构体
	cat11 := <-carChan
	cat22 := <-carChan
	fmt.Println(cat11, cat22)

	//4: 结构体指针
	var catPointChan chan *Cat
	catPointChan = make(chan *Cat, 10)
	catPointChan <- &cat1
	catPointChan <- &cat2

	//5: 存放任意数据类型
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)

	allChan <- cat1
	allChan <- cat2
	allChan <- num1
	allChan <- num2
	allChan <- m1
	allChan <- m2

}

type Cat struct {
	Name string
	Age  int
}
