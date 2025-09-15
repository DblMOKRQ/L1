package main

import (
	"fmt"
)

func main() {
	temperature := []float64{0.0, 0.9, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	res := make(map[int][]float64)
	for _, temp := range temperature {
		decade := int(temp) / 10 * 10
		res[decade] = append(res[decade], temp)
	}
	for k, v := range res {
		fmt.Println(k, v)
	}
}
