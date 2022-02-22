package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type MytypeInt int

const (
	SidesTriangle MytypeInt = 3
	SidesSquare   MytypeInt = 4
	SidesCircle   MytypeInt = 0
)

func CalcSquare(sideLen float64, sidesNum MytypeInt) float64 {
	var result float64

	switch sidesNum {
	case SidesTriangle:
		result = (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
	case SidesSquare:
		result = math.Pow(sideLen, 2)
	case SidesCircle:
		result = math.Pi * math.Pow(sideLen, 2)
	}
	return result
}
