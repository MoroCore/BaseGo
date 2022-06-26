package main

import (
	"fmt"
	"runtime"
)

/**
 * @GoName demo02
 * @Author Crow
 * @Email 648960069@qq.com
 * @Date 2022/6/25 17:46
 */
func main() {

	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu)
	fmt.Print("cpu : ", cpu)
}
