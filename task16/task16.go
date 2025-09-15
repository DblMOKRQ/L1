package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var left, right []int
	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}
func main() {
	arr := []int{7, 2, 1, 6, 8, 5, 3, 4}
	fmt.Println(quickSort(arr))
}
