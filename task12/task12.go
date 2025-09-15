package main

import (
	"fmt"
	"sort"
)

func sol(arr []string) []string {
	if len(arr) == 0 {
		return []string{}
	}
	var res []string
	sort.Strings(arr)
	l := 0
	res = append(res, arr[l])
	for r := 1; r < len(arr); r++ {
		if arr[l] != arr[r] {
			res = append(res, arr[r])
			l = r
		}
	}

	return res
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(sol(arr))
}
