// Package hamming provides functionality to calculate the Hamming distance between two DNA strands.
package hamming

import (
	"errors"
	"regexp"
)

// Distance calculates the Hamming distance between two DNA strands.
// It returns an error if the strands are not of the same length or if they contain invalid nucleotide characters.
// The Hamming distance is computed by comparing each nucleotide in the two strands and counting the mismatches.
//
// Parameters:
//
//	a (string): The first DNA strand.
//	b (string): The second DNA strand.
//
// Returns:
//
//	int: The Hamming distance between the two DNA strands.
//	error: An error is returned if the strands are not of equal length or contain invalid characters.
func Distance(a, b string) (int, error) {
	// Ensure the strands are of equal length
	if len(a) != len(b) {
		return 0, errors.New("strings must be of equal length")
	}

	// Validate that both strands consist only of valid nucleotides (C, A, G, T)
	if !isValidStrand(a) || !isValidStrand(b) {
		return 0, errors.New("strands must be of valid nucleotide sequence")
	}

	dis := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dis++
		}
	}

	return dis, nil
}

// isValidStrand checks if a given DNA strand contains only valid nucleotides (C, A, G, T).
// The strand must contain at least one character and only valid nucleotide characters.
//
// Parameters:
//
//	s (string): The DNA strand to be validated.
//
// Returns:
//
//	bool: Returns true if the strand contains only valid nucleotides (C, A, G, T), otherwise false.
func isValidStrand(s string) bool {
	re := regexp.MustCompile("^[CAGT]*$") 
	return re.MatchString(s)
}
