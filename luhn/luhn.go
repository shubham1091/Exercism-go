package luhn

import (
	"unicode"
)

// Valid determines if a string is a valid number per the Luhn algorithm, commonly used for verifying credit card and other identification numbers.
// It trims whitespace, checks length, and iterates over each character from right to left to calculate a sum based on even or odd positions.
// Returns true if the sum modulo 10 is zero, false otherwise.
func Valid(num string) bool {
	sum, count := 0, 0

	for i := len(num) - 1; i >= 0; i-- {
		ch := rune(num[i])
		switch {
		case ch == ' ':
			continue
		case unicode.IsDigit(ch):
			digit := int(ch - '0')
			if count%2 == 1 {
				digit <<= 1
			}
			if digit > 9 {
				digit -= 9
			}
			sum += digit
			count++
		default:
			return false
		}
	}

	return count>1 &&  sum%10 == 0

}
