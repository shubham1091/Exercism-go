package darts

import "math"

// insideCircle checks if the point (x, y) is within the circle of the given radius.
func insideCircle(x, y, radius float64) bool {
	return (x*x + y*y) <= radius*radius
}

// RadiusPoints associates a circle's radius with the points scored for landing within it.
type RadiusPoints struct {
	Radius float64
	Points int
}

// Define constants for the radii of the scoring circles.
const (
	innerRadius  = 1.0
	middleRadius = 5.0
	outerRadius  = 10.0
)

// scoringTable maps radii to their corresponding points.
var scoringTable = []RadiusPoints{
	{Radius: innerRadius, Points: 10},
	{Radius: middleRadius, Points: 5},
	{Radius: outerRadius, Points: 1},
	{Radius: math.Inf(1), Points: 0}, // Points outside the outer circle
}

// Score calculates the score for a dart landing at coordinates (x, y).
func Score(x, y float64) int {
	for _, entry := range scoringTable {
		if insideCircle(x, y, entry.Radius) {
			return entry.Points
		}
	}
	return 0
}
