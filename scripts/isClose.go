package scripts

import "math"

// check if two objects are close
func IsClose(x1, y1, x2, y2, maxDistance float64) bool {
	dx := x2 -x1
	dy := y2 -y1
	distance := math.Sqrt(dx*dx +dy*dy)
	return distance < maxDistance
}