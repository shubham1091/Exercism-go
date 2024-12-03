// Package proverb provides functionality to generate a proverb based on a list of inputs.
package proverb

import "fmt"

// Proverb generates a proverb from a list of inputs.
// It returns a slice of strings, each representing a line of the proverb.
func Proverb(rhyme []string) []string {
	lineCount := len(rhyme)

	result := make([]string, 0, lineCount)

	if lineCount == 0 {
		return result
	}

	for idx, word := range rhyme[1:] {
		line := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[idx], word)
		result = append(result, line)
	}
	finalLine := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	result = append(result, finalLine)

	return result
}
