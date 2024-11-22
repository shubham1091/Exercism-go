package bob

import (
	"strings"
	"unicode"
)

// Remark represents a user input (or remark) that can be analyzed for specific properties.
type Remark struct {
	remark string // The trimmed content of the user's input.
}

// newRemark initializes a new Remark by trimming any leading or trailing whitespace.
//
// Parameters:
//   - remark: A string containing the user's input.
//
// Returns:
//   - A Remark struct with the trimmed input.
func newRemark(remark string) Remark {
	return Remark{strings.TrimSpace(remark)}
}

// String returns the raw string content of the Remark.
//
// Returns:
//   - The string value of the remark.
func (r *Remark) String() string {
	return r.remark
}

// IsEmpty checks if the remark is empty or contains only whitespace.
//
// Returns:
//   - true if the remark is empty or consists only of whitespace; otherwise, false.
func (r *Remark) IsEmpty() bool {
	return r.remark == ""
}

// IsQuestion checks if the remark ends with a question mark.
//
// Returns:
//   - true if the remark ends with a '?'; otherwise, false.
func (r *Remark) IsQuestion() bool {
	return strings.HasSuffix(r.remark, "?")
}

// IsYelling checks if the remark is considered yelling.
// Yelling is defined as containing letters, with all letters being uppercase.
//
// Returns:
//   - true if the remark contains at least one letter and all letters are uppercase.
//   - false if there are no letters or any lowercase letters are present.
func (r *Remark) IsYelling() bool {
	hasLetters := strings.IndexFunc(r.remark, unicode.IsLetter) != -1
	isUppercased := strings.ToUpper(r.remark) == r.remark
	return hasLetters && isUppercased
}

// IsExasperated checks if the remark is both yelling and a question.
// A remark is exasperated if it is yelling (all letters uppercase) and ends with a question mark.
//
// Returns:
//   - true if the remark is both yelling and a question; otherwise, false.
func (r *Remark) IsExasperated() bool {
	return r.IsQuestion() && r.IsYelling()
}
