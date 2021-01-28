package main

import (
	"fmt"
)

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxContiguousArrayLinear(arr []int) int {
	var (
		size                 = len(arr) - 1
		currentMaxNumber     = arr[0]
		singleNumberMaxArray = arr[0]
	)

	for i := 1; i < size; i++ {
		singleNumberMaxArray = max(singleNumberMaxArray+arr[i], arr[i])
		currentMaxNumber = max(currentMaxNumber, singleNumberMaxArray)
	}

	return currentMaxNumber
}

func main() {
	var (
		arr = []int{-9, 0, 3, 7, -3, -6, -8, 9, 1, -5, 8, -1, -4, 6, 6, 1}
	)
	x1 := maxContiguousArrayLinear(arr)
	fmt.Println(x1)
}
