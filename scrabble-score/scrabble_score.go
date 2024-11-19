// Package scrabble provides a function to calculate the Scrabble score of a given word.
//
// The score of a word is determined by summing the individual letter values.
// The values for each letter are based on the official Scrabble letter distribution.
package scrabble

import "strings"

// Score calculates the Scrabble score of the given word.
//
// The function takes a word as input, converts it to lowercase, and then calculates the score by summing the values of its letters.
// The value of each letter is determined based on the official Scrabble letter distribution:
// A, E, I, O, U, L, N, R, S, T: 1 point
// D, G: 2 points
// B, C, M, P: 3 points
// F, H, V, W, Y: 4 points
// K: 5 points
// J, X: 8 points
// Q, Z: 10 points
//
// The function returns the total score of the word.
func Score(word string) int {
	word = strings.ToLower(word) // Convert the word to lowercase to handle case insensitivity
	score := 0

	// Iterate through each letter in the word and add the appropriate score
	for _, letter := range word {
		switch letter {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score += 1
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		default:
			score += 0 // Ignore non-letter characters
		}
	}
	return score
}
