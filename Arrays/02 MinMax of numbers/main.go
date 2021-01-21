package main

import "fmt"

func minmaxArrayLinearSearch(arr []int, min, max, size int) (int, int) {
	if size == 1 {
		min = arr[0]
		max = arr[0]
		return min, max
	} else if arr[0] > arr[1] {
		max = arr[0]
		min = arr[1]
	}

	for i := 2; i < size; i++ {
		if arr[i] > max {
			max = arr[i]
		} else if arr[i] < min {
			min = arr[i]
		}
	}
	return min, max
}

func main() {
	var (
		arr = []int{23, 21, 36, 27, 6, 34, 12}
		min = arr[0]
		max = arr[1]
	)
	fmt.Println("Original array :-", arr)
	min, max = minmaxArrayLinearSearch(arr, min, max, len(arr))
	fmt.Println("Smallest number in array is :-", min, "\nLargest number in array is :-", max)
}
