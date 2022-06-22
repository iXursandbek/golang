package main

import (
	"fmt"
	"math"
)

func main() {
	var len float64
	var num intCustomType
	fmt.Scan(&len, &num)
	fmt.Println(CalcSquare(len, intCustomType(num)))
}

type intCustomType struct{
	SidesTriangle int
	SidesSquare
	SidesCircle
}

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	var s float64
	if sidesNum == 3 {
		s = sideLen * sideLen * math.Sqrt(3) / 4
	} else if sidesNum == 4 {
		s = sideLen * sideLen
	} else if sidesNum == 0 {
		s = math.Pi * sideLen * sideLen
	}
	return s
}
