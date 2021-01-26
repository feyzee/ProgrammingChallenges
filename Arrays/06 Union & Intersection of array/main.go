package main

import "fmt"

func union(arr1, arr2 []int, m, n int) []int {
	var tmp []int
	i, j := 0, 0

	for i < m && j < n {
		if arr1[i] < arr2[j] {
			fmt.Println(arr1[i])
			i++
		} else if arr1[i] > arr2[j] {
			fmt.Println(arr2[j])
			j++
		} else {
			fmt.Println(arr1[i])
			i++
			j++
		}
	}

	for i < m {
		fmt.Println(arr1[i])
		i++
	}

	for j < n {
		fmt.Println(arr2[j])
		j++
	}

	return tmp
}

func intersection(arr1, arr2 []int, m, n int) []int {
	var tmp []int
	fmt.Println(arr1, arr2)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr1[i] == arr2[j] {
				fmt.Println(arr1[i])
			}
		}
	}

	return tmp
}

func main() {
	var (
		arr1 = []int{2, 5, 6, 9}
		arr2 = []int{4, 5, 8, 10}
		m    = len(arr1)
		n    = len(arr2)
	)
	fmt.Println(arr1, arr2)
	unionArray := union(arr1, arr2, m, n)
	intersectionArray := intersection(arr1, arr2, m, n)
	fmt.Println(unionArray)
	fmt.Println(intersectionArray)
}
