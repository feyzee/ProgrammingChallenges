package main

import "fmt"

func intersection(arr1, arr2 []int) []int {
	var (
		tmp []int
		m   = len(arr1)
		n   = len(arr2)
	)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr1[i] == arr2[j] {
				tmp = append(tmp, arr1[i])
			}
		}
	}

	return tmp
}

func commonElementsArray(arr1, arr2, arr3 []int) []int {
	// This method uses a function -- intersection() -- to find
	// intersection of the array.
	x := intersection(arr1, arr2)
	y := intersection(x, arr3)
	return y
}

func commonElementsArray2(arr1, arr2, arr3 []int) []int {
	// This method uses a simple for loop to find
	// intersection of the array.
	var (
		i, j, k = 0, 0, 0
		arr     []int
	)

	for len(arr1) > i && len(arr2) > j && len(arr3) > k {
		if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
			arr = append(arr, arr1[i])
			i++
			j++
			k++
		} else if arr1[i] < arr2[j] {
			i++
		} else if arr2[j] < arr3[k] {
			j++
		} else {
			k++
		}
	}
	return arr
}

func main() {
	var (
		arr1 = []int{1, 5, 10, 20, 40, 80}
		arr2 = []int{6, 7, 20, 80, 100}
		arr3 = []int{3, 4, 15, 20, 30, 70, 80, 120}
	)

	x := commonElementsArray(arr1, arr2, arr3)
	y := commonElementsArray2(arr1, arr2, arr3)
	fmt.Println(x, y)
}
