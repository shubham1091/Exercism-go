// Package triangle provides functionality to classify triangles based on the lengths of their sides.
package triangle

// Kind represents the type of a triangle.
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides determines the type of a triangle based on the lengths of its sides.
// It takes three float64 values representing the lengths of the sides.
// It returns a Kind value indicating the type of the triangle (NaT, Equ, Iso, or Sca).
func KindFromSides(a, b, c float64) Kind {
	if isTriangle(a, b, c) {
		switch {
		case a == b && b == c:
			return Equ
		case a == b || b == c || a == c:
			return Iso
		default:
			return Sca
		}
	} else {
		return NaT
	}
}

func isTriangle(a, b, c float64) bool {
	return a > 0 && b > 0 && c > 0 &&
		(a+b >= c) && (b+c >= a) && (a+c >= b)
}
