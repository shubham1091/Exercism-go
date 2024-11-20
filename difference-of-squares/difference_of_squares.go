// Package diffsquares is a small library for calculating values relating to squares of a number.
package diffsquares

// formulae found at
// https://learnersbucket.com/examples/algorithms/difference-between-square-of-sum-of-numbers-and-sum-of-square-of-numbers/

// SquareOfSum returns the square of the sum of numbers from 1 to the number passed in.
func SquareOfSum(num int) int {
	sum := (num * (num + 1)) / 2
	return sum * sum
}

// SumOfSquares returns the sum of the squared numbers from 1 to the number passed in.
func SumOfSquares(num int) int {
	return (num * (num + 1) * ((num * 2) + 1)) / 6
}

// Difference returns the square of sums minus the sum of squares for the number passed.
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}