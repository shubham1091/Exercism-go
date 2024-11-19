package raindrops

import (
	"strconv"
	"strings"
)

// drop represents a number divisor and its associated sound.
type drop struct {
	divisor int
	sound   string
}

// Convert takes an integer and returns a string based on its divisibility by 3, 5, or 7.
// If divisible by 3, it adds "Pling"; if divisible by 5, it adds "Plang";
// if divisible by 7, it adds "Plong". If none, it returns the number as a string.
func Convert(number int) string {
	var sound strings.Builder

	drops := []drop{
		{divisor: 3, sound: "Pling"},
		{divisor: 5, sound: "Plang"},
		{divisor: 7, sound: "Plong"},
	}
	for _, dp := range drops {
		if number%dp.divisor == 0 {
			sound.WriteString(dp.sound)
		}
	}
	if sound.Len() == 0 {
		return strconv.Itoa(number)
	}
	return sound.String()
}
