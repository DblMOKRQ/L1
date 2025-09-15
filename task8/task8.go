package main

import "fmt"

func setBitTo(n int64, i uint, bit int) int64 {
	if bit != 0 && bit != 1 {
		return n
	}
	mask := int64(1) << i
	if bit == 1 {
		return n | mask
	}
	return n &^ mask
}

func main() {
	n := int64(5)

	orig := n
	res := setBitTo(orig, 2, 0)

	fmt.Printf("original: %d (0b%064b)\n", orig, uint64(orig))
	fmt.Printf("set bit %d -> %d: %d (0b%064b)\n", 1, 0, res, uint64(res))

}
