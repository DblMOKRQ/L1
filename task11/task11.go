package main

import (
	"fmt"
	"sort"
)

func sol(a []int, b []int) []int {
	sort.Ints(a)
	sort.Ints(b)
	l := 0
	r := 0
	res := make([]int, 0)
	for l < len(a) && r < len(b) {
		if a[l] == b[r] {
			res = append(res, a[l])
			l++
			r++
		} else if a[l] < b[r] {
			l++
		} else if a[l] > b[r] {
			r++
		}
	}
	return res
}
func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	fmt.Println(sol(a, b))
}
