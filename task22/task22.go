package main

import (
	"fmt"
	"math/big"
)

func multi(x *big.Float, y *big.Float) *big.Float {
	return big.NewFloat(0).Mul(x, y)
}

func quo(x *big.Float, y *big.Float) *big.Float {
	return big.NewFloat(0).Quo(x, y)
}

func plus(x *big.Float, y *big.Float) *big.Float {
	return big.NewFloat(0).Add(x, y)
}

func minus(x *big.Float, y *big.Float) *big.Float {
	return big.NewFloat(0).Sub(x, y)
}

func main() {
	x := big.NewFloat(1500000)
	y := big.NewFloat(2000000)

	fmt.Printf("x = %v, y = %v\n", x, y)
	fmt.Printf("x + y = %v\n", plus(x, y))
	fmt.Printf("x - y = %v\n", minus(x, y))
	fmt.Printf("x * y = %v\n", multi(x, y))
	fmt.Printf("x / y = %.10f\n", quo(x, y))

	fmt.Printf("x после операций = %v\n", x)
	
}
