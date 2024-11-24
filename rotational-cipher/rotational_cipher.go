package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	shiftKey = shiftKey % 26

	key := rune(shiftKey)

	str := strings.Builder{}
	for _, char := range plain {
		switch {
		case unicode.IsUpper(char):
			rotated := 'A' + (char-'A'+key)%26
			str.WriteRune(rotated)
		case unicode.IsLower(char):
			rotated := 'a' + (char-'a'+key)%26
			str.WriteRune(rotated)
		default:
			str.WriteRune(char)
		}
	}
	return str.String()
}
