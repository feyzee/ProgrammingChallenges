package main

import "fmt"

func cyclicRotateArray(arr []int) []int {
	i := len(arr)
	x := arr[i-1]
	arr = append([]int{x}, arr...)
	arr = arr[:i]
	return arr
}

func main() {
	var (
		arr = []int{9, 0, 1, 2, 3, 6, 6, 6, 4, 5, 4}
	)
	arr = cyclicRotateArray(arr)
	fmt.Println(arr)
}
