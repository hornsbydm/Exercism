package triangle

import "math"

// Kind describes a type of triangle.
type Kind int

const (
	// NaT is Not a Triangle
	NaT Kind = iota
	//Equ is an equilateral triangle (i.e. a=b=c)
	Equ
	//Iso is an isoceles triangle
	Iso
	//Sca is a scalene triangle.
	Sca
)

//KindFromSides determines the kind of triangle given the side lengths.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case math.IsNaN(a), math.IsNaN(b), math.IsNaN(c):
		return NaT
	case math.IsInf(a, 0), math.IsInf(b, 0), math.IsInf(c, 0):
		return NaT
	case a <= 0, b <= 0, c <= 0:
		return NaT
	case a+b < c, c+a < b, b+c < a:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b, a == c, b == c:
		return Iso
	}
	return Sca
}
