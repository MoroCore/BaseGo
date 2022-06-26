package main

import (
	"fmt"
	"time"
)

/*
 * @GoName main
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022 18:07
 */

func main() {

	count := 0
	startTime := time.Now().Unix()
	var flag bool
	for j := 2; j < 1000000; j++ {
		for i := 2; i < j; i++ {
			if j%i == 0 {
				flag = false
				break
			}
			if flag {
				count++
			}
		}
	}
	endTime := time.Now().Unix()
	fmt.Println("spend time = ", endTime-startTime)
}
