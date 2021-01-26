package main

import "fmt"

func sort012(arr []int) []int {
	arraySize := len(arr)
	low, mid, high := 0, 0, arraySize-1
	for mid <= high {
		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}

	return arr
}

func main() {
	var (
		arr = []int{0, 1, 2, 0, 0, 1, 2, 0, 1, 2}
	)
	fmt.Println("Original array -", arr)
	sort012(arr)
	fmt.Println("Sorted array -", arr)
}
