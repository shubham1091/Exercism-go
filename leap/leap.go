package leap

// IsLeapYear determines if the given year is a leap year.
//
// A year is a leap year if:
//   - It is divisible by 400, or
//   - It is divisible by 4 but not divisible by 100.
//
// The function returns true if the year is a leap year, and false otherwise.
func IsLeapYear(year int) bool {
	if year < 0 {
		return false
	}
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	case year%4 == 0:
		return true
	default:
		return false
	}
}
