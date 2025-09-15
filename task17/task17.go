package main

import "fmt"

func binarySearch(arr []int, targer int) int {
	curIndex := (len(arr) - 1) / 2
	left := 0
	right := len(arr) - 1
	for arr[curIndex] != targer && left <= right {
		if arr[curIndex] > targer {
			right = curIndex - 1
		} else {
			left = curIndex + 1
		}
		curIndex = (left + right) / 2
	}
	if arr[curIndex] == targer {
		return curIndex
	}
	return -1
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(arr, 9))
}
