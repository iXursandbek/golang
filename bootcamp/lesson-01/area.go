package main

import (
	"fmt"
	"math"
)

func main() {
	var R float32
	fmt.Scan(&R)

	fmt.Printf("Area: %0.2f\n", area(R))
}

func area(r float32) (area1 float32) {
	return r * r * (4 - math.Pi)
}
