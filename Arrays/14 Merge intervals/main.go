package main

import (
	"fmt"
	"sort"
)

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func mergeIntervals(array [][]int) [][]int {
	mergedArray := make([][]int, 0)

	sort.Slice(array, func(x, y int) bool { return array[x][0] < array[y][0] })

	mergedArray = append(mergedArray, array[0])

	for i := 1; i < len(array); i++ {
		if array[i][0] <= mergedArray[len(mergedArray)-1][1] {
			mergedArray[len(mergedArray)-1][1] = max(mergedArray[len(mergedArray)-1][1], array[i][1])
		} else {
			mergedArray = append(mergedArray, array[i])
		}
	}

	return mergedArray
}

func main() {
	arr := [][]int{{1, 2}, {2, 4}, {5, 8}, {9, 14}}
	x := mergeIntervals(arr)
	fmt.Println(x)
}
