package resistorcolorduo

var bands = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	// Handle empty or insufficient colors
    if len(colors) < 2 {
        return 0 // or handle error according to requirements
    }

    // Calculate resistance using only first two colors
    firstDigit, exists1 := bands[colors[0]]
    secondDigit, exists2 := bands[colors[1]]

    // Handle invalid colors
    if !exists1 || !exists2 {
        return 0 // or handle error according to requirements
    }

    return firstDigit*10 + secondDigit
}
