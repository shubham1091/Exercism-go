// Package romannumerals provides functionality to convert integers to Roman numerals.
package romannumerals

import (
	"errors"
	"strings"
)

// RomanNumeral represents a mapping between an integer value and its Roman numeral representation.
type RomanNumeral struct {
	value int
	roman string
}

// Predefined mapping of integer values to Roman numerals, sorted in descending order.
var romanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral converts an integer to its Roman numeral representation.
// Parameters:
//   - input: An integer between 1 and 3999.
//
// Returns:
//   - A string representing the Roman numeral equivalent of the input.
//   - An error if the input is outside the valid range.
func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("input must be between 1 and 3999")
	}

	var builder strings.Builder

	for _, numeral := range romanNumerals {
		for input >= numeral.value {
			builder.WriteString(numeral.roman)
			input -= numeral.value
		}
	}

	return builder.String(), nil
}
