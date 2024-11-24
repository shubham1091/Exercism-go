package isbn

import "strings"

// IsValidISBN validates if a given string is a valid ISBN-10.
func IsValidISBN(isbn string) bool {
	// Remove hyphens and trim surrounding whitespace
	isbn = strings.ReplaceAll(strings.TrimSpace(isbn), "-", "")

	// ISBN-10 must have exactly 10 characters after removing hyphens
	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for idx, char := range isbn {
		if idx == 9 && char == 'X' { // 'X' is valid only in the last position and represents 10
			sum += 10
			continue
		}

		if char < '0' || char > '9' { // Ensure all other characters are numeric
			return false
		}

		// Multiply digit by its weight (10 - idx) and add to the sum
		sum += int(char-'0') * (10 - idx)
	}

	// Valid ISBN-10 must have a checksum divisible by 11
	return sum%11 == 0
}
