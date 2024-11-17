package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber returns a formatted string description of a number
// with exactly one decimal place.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

// NumberBox interface defines a container that holds a number
type NumberBox interface {
	Number() int
}

// DescribeNumberBox formats the contained number as a float with one decimal place
func DescribeNumberBox(nb NumberBox) string {
	if nb == nil {
		return "This is a box containing the number 0.0"
	}
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

// FancyNumber represents a number stored as a string
type FancyNumber struct {
	n string
}

// Value returns the string representation of the fancy number
func (i FancyNumber) Value() string {
	return i.n
}

// FancyNumberBox interface defines a container that holds a string value
type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber safely extracts an integer from a FancyNumberBox
func ExtractFancyNumber(fnb FancyNumberBox) int {
	// Handle nil input
	if fnb == nil {
		return 0
	}

	// Type assertion to check if it's specifically a FancyNumber
	fancyNum, ok := fnb.(FancyNumber)
	if !ok {
		return 0
	}

	// Get the value and ensure it's not empty
	value := fancyNum.Value()
	if value == "" {
		return 0
	}

	// Convert string to integer, handling any conversion errors
	num, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return num
}

// DescribeFancyNumberBox returns a formatted description of a FancyNumberBox
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	num := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(num))
}

// DescribeAnything returns a description of the provided value based on its type
func DescribeAnything(i interface{}) string {
	// Handle nil input
	if i == nil {
		return "Return to sender"
	}

	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	default:
		return "Return to sender"
	}
}
