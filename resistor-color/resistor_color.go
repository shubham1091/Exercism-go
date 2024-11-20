package resistorcolor


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

var colorList []string

func init() {
	for color := range bands {
		colorList = append(colorList, color)
	}
}

// Colors returns the list of all colors sorted alphabetically.
func Colors() []string {
	return colorList
}

// ColorCode returns the resistance value of the given color.
// Returns -1 if the color is not found to indicate an error.
func ColorCode(color string) int {
	if code, exists := bands[color]; exists {
		return code
	}
	return -1 // Return -1 to indicate the color was not found
}
