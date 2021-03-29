package test

const (
	EPSILON float64 = 0.00000001
)

func FloatEquals(left float64, right float64) bool {
	return (left-right) < EPSILON && (right-left) < EPSILON
}
