package grains

import (
	"errors"
)

// Square calculates the number of grains on a specific square of the chessboard.
// Parameters:
//   - number: The square number (1 through 64).
//
// Returns:
//   - The number of grains on the given square.
//   - An error if the square number is invalid.
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("invalid square number")
	}
	// Use bit-shifting for fast power of 2 computation: 2^(number-1)
	return 1 << (number - 1), nil
}

// Total calculates the total number of grains on the chessboard.
// Returns:
//   - The total number of grains (using the formula for the sum of a geometric series).
func Total() uint64 {
	// Sum of a geometric series: 2^64 - 1
	return 1<<64 - 1
}
