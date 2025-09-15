package main

import "fmt"

func main() {
	s := "главрыба🫥"
	runes := []rune(s)
	l := 0
	r := len(runes) - 1
	for l < r {
		temp := runes[l]
		runes[l] = runes[r]
		runes[r] = temp
		l++
		r--
	}
	fmt.Println(string(runes))
}
