package main

import (
	"fmt"
)

func main() {

	//1: map的声明
	// var mapList map[int]int

	//map的声明 举例
	var a map[string]string
	// var b map[string]int
	// var c map[int]string
	// var d map[string]map[string]string
	// 注意:声明是不会分配内存的，初始化需要make ,分配内存后才能赋值和使用。

	a = make(map[string]string, 10)
	a["1"] = "宋江"
	a["2"] = "吴用"
	// 1) map在使用前一定要make
	// 2) map的 key是不能重复，如果重复了，则以最后这个key-value为准
	// 3) map的value是可以相同的.
	// 4) map的 key-value是无序
	// 5) make内置函数数目

	//2 map的使用
	//方式1  先定义再分配空间
	var b map[string]string
	b = make(map[string]string)
	b["01"] = "moro"
	b["02"] = "moro"
	fmt.Println("---------------------------------")
	fmt.Println(b["01"])
	//方式2: 短声明
	c := make(map[string]string)
	c["01"] = "moro"
	//方式3: 短声明 并初始化参数
	heroes := map[string]string{
		"heri": "moro",
	}
	fmt.Printf(heroes["heri"])
	fmt.Println("heroes=", heroes)

	//我们要存放3个学生信息，每个学生有name和sex信息
	students := make(map[string]map[string]string)

	students["01"] = make(map[string]string)
	students["01"]["name"] = "moro"
	students["01"]["sex"] = "男"

	students["02"] = make(map[string]string)
	students["02"]["name"] = "carlos"
	students["02"]["age"] = "男"

	//map的增删改查
	//1 map的增加和更新
	// map[key] = value 如果key还没有，就是增加，如果key存在就是修改。

	//2 map的删除
	// delete(map，"key"),delete是一个内置函数，如果key存在，就删除该key-value,如果key不存在，不操作，但是也不会报错
	// map 切片  通道的 没有分配空间的空值是 nil
	// 演示delete
	delete(students, "01")
	// 如果想要删除map全部的元素  2种方式  遍历删除   重新分配空间
	value, ok := c["01"]
	if ok {
		fmt.Printf("有no1 key 值为%vln", value)
	} else {

	}
	// 说明:如果heroes这个map中存在"no1"，那么findRes就会返回true,否则返回false
	//4 map的遍历
	fmt.Println()
	for k1, v1 := range students {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v \t", k2, v2)
		}
		fmt.Println()
	}
}
