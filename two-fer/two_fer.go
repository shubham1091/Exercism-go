// Package twofer provides a function to share cookies in a "two-fer" offer.
package twofer

import (
	"fmt"
	"strings"
)

// ShareWith returns a phrase indicating who gets the extra cookie.
// If the name is empty, it defaults to "you".
func ShareWith(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
