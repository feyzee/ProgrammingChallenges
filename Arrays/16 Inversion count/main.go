package main

import "fmt"

func countInversion(arr []int) int {
	count := 0

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[i] > arr[j] {
				count++
			}
		}
	}

	return count
}

func main() {
	arr := []int{2, 4, 1, 3, 5}
	count := countInversion(arr)
	fmt.Println(count)
}
