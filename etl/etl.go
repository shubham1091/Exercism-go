package etl

import "strings"

// Transform takes a map where the key is the score and the value is a list of letters
// and returns a map where each letter is mapped to its corresponding score.
func Transform(in map[int][]string) map[string]int {
	// Create an empty map to store the transformed letter-score pairs.
	out := make(map[string]int)

	// Iterate over each score and its corresponding list of letters.
	for score, letters := range in {
		// For each letter in the list, map it to its score (after converting it to lowercase).
		for _, letter := range letters {
			out[strings.ToLower(letter)] = score
		}
	}

	// Return the transformed map.
	return out
}
