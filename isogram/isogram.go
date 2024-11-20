package isogram

import "strings"

// IsIsogram determines if a word or phrase is an isogram, ignoring spaces and hyphens.
func IsIsogram(word string) bool {
	if word == "" {
		return true // An empty string is a trivial isogram
	}
	word = strings.ToLower(word) 

	seenChars := make(map[rune]bool)

	for _, letter := range word {
		if letter != ' ' && letter != '-' {
			if seenChars[letter] {
				return false // Duplicate character found, not an isogram
			}
			seenChars[letter] = true // Mark this letter as seen
		}
	}

	return true 
}
