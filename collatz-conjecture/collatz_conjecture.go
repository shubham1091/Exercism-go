package collatzconjecture

import "errors"

// CollatzConjecture takes a positive integer n and returns the number of steps
// required to reach 1 based on the Collatz Conjecture.
// If n is not positive, it returns an error.
func CollatzConjecture(n int) (int, error) {
	// Validate that n is a positive integer
	if n <= 0 {
		return 0, errors.New("n must be a positive integer greater than 0")
	}

	steps := 0
	
	// Continue the process until n becomes 1
	for n > 1 {
		if n%2 == 0 {
			n >>=1
		} else {
			n = n*3 + 1
		}
		steps++
	}
	return steps, nil
}
