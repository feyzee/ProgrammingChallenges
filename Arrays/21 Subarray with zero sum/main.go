package main

import "fmt"

func subarrayZeroSum(arr []int) bool {
	for i := range arr {
		currentSum := arr[i]
		j := i + 1

		for j <= len(arr) {
			if currentSum == 0 {
				return true
			} else if j == len(arr) {
				break
			}

			currentSum = currentSum + arr[j]
			j++
		}
	}
	return false
}

func main() {
	arr := []int{4, 2, -3, 1, 6}
	// arr := []int{-3, 2, 3, 1, 6}

	if subarrayZeroSum(arr) == true {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}
