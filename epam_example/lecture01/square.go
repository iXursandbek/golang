package lecture01

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type Sides int

const (
	SidesCircle   Sides = 0
	SidesTriangle Sides = 3
	SidesSquare   Sides = 4
)


func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	var s float64
	if sidesNum == SidesTriangle {
		s = sideLen * sideLen * math.Sqrt(3) / 4
	} else if sidesNum == SidesSquare {
		s = sideLen * sideLen
	} else if sidesNum == SidesCircle {
		s = math.Pi * sideLen * sideLen
	} else {
		s = 0
	}
	return s

}

func main() {
	fmt.Println(CalcSquare(10.0, SidesCircle))
}
