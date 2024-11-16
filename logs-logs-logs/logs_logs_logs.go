package logs

import (
	"strings"
	"unicode/utf8"
)

// Define constants for application names.
const (
	recommendationApp = "recommendation"
	searchApp        = "search"
	weatherApp       = "weather"
	defaultApp       = "default"
)

// Mapping of special characters to application names
var rep = map[rune]string{
	'❗': recommendationApp,
	'🔍': searchApp,
	'☀': weatherApp,
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	// Iterate through the characters of the log string.
	for _, c := range log {
		if app, ok := rep[c]; ok {
			return app
		}
	}
	// Return default if no known character is found.
	return defaultApp
}

// Replace replaces all occurrences of the corrupted character with the original character.
func Replace(log string, corrupted, original rune) string {
	// Directly replace runes using strings.ReplaceAll
	return strings.ReplaceAll(log, string(corrupted), string(original))
}

// WithinLimit determines whether the log's character count is within the limit.
func WithinLimit(log string, limit int) bool {
	// Count the number of runes (characters) in the log.
	return utf8.RuneCountInString(log) <= limit
}
