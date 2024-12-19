package anagram

import (
	"strings"
	"unicode"
)

// Detect function detects anagrams in a list of candidates for the given subject
func Detect(subject string, candidates []string) []string {
	var result []string
	normalizedSubject := strings.ToLower(subject) // Normalize the subject to lowercase

	// Iterate over the candidate words
	for _, candidate := range candidates {
		// Skip the word if it's the same as the subject (case-insensitive)
		if strings.ToLower(candidate) == normalizedSubject {
			continue
		}

		// If the candidate is an anagram of the subject, add it to the result
		if isAnagram(normalizedSubject, candidate) {
			result = append(result, candidate)
		}
	}
	return result
}

// isAnagram returns true if two strings are anagrams of each other
func isAnagram(s1, s2 string) bool {
	// If the lengths are different, they can't be anagrams
	if len(s1) != len(s2) {
		return false
	}

	// Create frequency maps for both strings
	s1Map := make(map[rune]int)
	s2Map := make(map[rune]int)

	// Count frequency of characters in the first string
	for _, r := range s1 {
		s1Map[rune(unicode.ToLower(r))]++
	}

	// Count frequency of characters in the second string
	for _, r := range s2 {
		s2Map[rune(unicode.ToLower(r))]++
	}

	// Compare the frequency maps
	for k, v := range s1Map {
		if s2Map[k] != v {
			return false
		}
	}

	return true
}
