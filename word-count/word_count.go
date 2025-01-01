package wordcount

import (
	"regexp"
	"strings"
)

// Frequency represents the mapping of words to their counts.
type Frequency map[string]int

// WordCount counts the frequency of each word in a given phrase.
func WordCount(phrase string) Frequency {
	freq := make(Frequency)

	// Convert the phrase to lowercase for case insensitivity.
	phrase = strings.ToLower(phrase)

	// Use a regular expression to match words (including contractions).
	// A word can contain alphanumeric characters and apostrophes (e.g., "that's").
	re := regexp.MustCompile(`[a-zA-Z0-9]+(?:'[a-zA-Z0-9]+)?`)

	// Find all matches in the phrase.
	words := re.FindAllString(phrase, -1)

	// Count each word.
	for _, word := range words {
		freq[word]++
	}

	return freq
}
