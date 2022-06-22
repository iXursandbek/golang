package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(MySquareRoot(10, 12))
	fmt.Println(MySquareRoot(4, 4))
}

func MySquareRoot(num, precision uint) (result float64) {
	var z float64
	var L10 float64 = float64(num)
	z = (1 + 28*L10 + 70*L10*L10 + 28*L10*L10*L10 + L10*L10*L10*L10) / (8 * (1 + L10) * (1 + 6*L10 + L10*L10))
	output := float64(math.Pow(10, float64(precision)))

	return float64(math.Round(z*output)) / output
}
