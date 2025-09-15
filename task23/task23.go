package main

import "fmt"

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	slice = slice[:len(slice)-1]
	return slice
}

func main() {
	s := []int{10, 20, 30, 40, 50}
	s = remove(s, 4)
	fmt.Println(s)
}
