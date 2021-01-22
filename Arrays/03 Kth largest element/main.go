package main

import (
	"fmt"
	"sort"
)

func kthNumber(arr []int, k int) (kth int) {
	// n := len(arr) - 1
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n-i-1; j++ {
	// 		temp := arr[i]

	// 	}
	// }
	sort.Ints(arr)
	kth = arr[k-1]
	return kth
}

func main() {
	var (
		arr = []int{23, 21, 36, 27, 6, 34, 12}
		k   = 5
	)
	fmt.Println("Original array :-", arr)
	fmt.Println("Value of K :-", k)
	kth := kthNumber(arr, k)
	fmt.Println("K'th number :-", kth)
}
