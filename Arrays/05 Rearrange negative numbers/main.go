package main

import "fmt"

func rearrange(arr []int) []int {
	var (
		size = len(arr)
		j    = 0
	)
	for i := 0; i < size; i++ {
		if arr[i] < 0 {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			j++
		}
	}
	return arr
}

func main() {
	var (
		arr = []int{-12, 11, -13, -5, 6, -7, 5, -3, -6}
	)
	x := rearrange(arr)
	fmt.Println(x)
}
