package main

import "fmt"

func sortTwoArraysOneSpace(arr1, arr2 []int) ([]int, []int) {
	m := len(arr1) - 1
	n := len(arr2) - 1

	for i := n; i >= 0; i-- {
		lastElement := arr1[m]
		j := arr1[m]

		for j = m - 1; j >= 0 && arr1[j] > arr2[i]; j-- {
			arr1[j+1] = arr1[j]
		}

		if j != m-1 || lastElement >= arr2[i] {
			arr1[j+1] = arr2[i]
			arr2[i] = lastElement
		}
	}

	return arr1, arr2
}

// func insertionSort(arr1 []int) {
// 	for i := 0; i < len(arr1)-1; i++ {
// 		current := arr1[i]
// 		j := i - 1

// 		for j >= 0 && arr1[j] > current {
// 			arr1[j+1] = arr1[j]
// 			j--
// 		}

// 		arr1[j+1] = current
// 	}
// }

func main() {
	var (
		arr1 = []int{1, 5, 9, 10, 15, 20}
		arr2 = []int{2, 3, 8, 13}
	)
	arr1, arr2 = sortTwoArraysOneSpace(arr1, arr2)
	fmt.Println(arr1, arr2)
}
