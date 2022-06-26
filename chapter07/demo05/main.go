package main

import "fmt"

func fnb(n int) []uint64 {

	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

//编写一个函数fbn(nint)，要求完成
func main() {
	fubSlice := fnb(1000)
	fmt.Println("fubSlice=", fubSlice)
}
