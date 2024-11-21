package pangram

import (
	"strings"
	"unicode"
)

const alphabetSize = 26

// IsPangram checks if the input string contains all letters of the alphabet
func IsPangram(input string) bool {
	input = strings.ToLower(input)
	seenLetters := make([]int, alphabetSize)

	lettersSeen := 0

	for _, letter := range input {
		// Skip non-alphabetic characters
		if !unicode.IsLetter(letter) {
			continue
		}

		// Calculate the index for the letter (a=0, b=1, etc.)
		index := int(letter - 'a')
		if seenLetters[index] == 0 {
			lettersSeen++
			seenLetters[index] = 1
		}

		// Early return if we've found all letters
		if lettersSeen == alphabetSize {
			return true
		}
	}

	return lettersSeen == alphabetSize
}
