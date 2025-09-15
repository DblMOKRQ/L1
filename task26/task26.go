package main

import (
	"fmt"
	"strings"
)

func uniqueChars(s string) bool {
	s = strings.ToLower(s)
	chars := make(map[rune]int)
	for _, v := range s {
		chars[v]++
	}
	for _, v := range chars {
		if v > 1 {
			return false
		}
	}
	return true
}

func main() {
	s := "aabcd"
	fmt.Println(uniqueChars(s))
}
