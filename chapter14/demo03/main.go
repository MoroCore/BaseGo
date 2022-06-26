package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @GoName demo03
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022/6/25 17:54
 */

var (
	myMap = make(map[int]int64, 10)
	lock  sync.Mutex
)

func main() {

	for i := 1; i <= 200; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 5)

	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d \n", i, v)
	}
	lock.Unlock()
}

func test(n int) {

	var res int64 = 1
	for i := 1; i <= int(n); i++ {
		res *= int64(i)
	}
	//加锁
	lock.Lock()
	myMap[n] = int64(res)
	lock.Unlock()
}
