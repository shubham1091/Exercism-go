package reverse

import "strings"

func Reverse(input string) string {
	var b strings.Builder

	str := []rune(input)

	for i := len(str) - 1; i >= 0; i-- {
		b.WriteRune(str[i])
	}

	return b.String()

}
