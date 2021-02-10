package main

import "fmt"

func countPairs(arr []int, sum int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				count++
			}
		}
	}
	return count
}

func main() {
	var (
		arr = []int{10, 12, 10, 15, -1, 7, 6, 5, 4, 2, 1, 1, 1}
		sum = 11
	)
	count := countPairs(arr, sum)
	fmt.Println("Number of pairs are - ", count)
}
