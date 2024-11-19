// Package gigasecond provides functionality to calculate the time that occurs
// exactly one gigasecond (1,000,000,000 seconds) after a given time.
// The purpose of this package is to demonstrate time manipulation using the
// time package in Go.
package gigasecond

import "time"

// AddGigasecond takes a time and returns the time that is exactly one gigasecond
// (1,000,000,000 seconds) later. The function works by adding a duration of one
// billion seconds to the given time using the Add method of time.Time.
func AddGigasecond(t time.Time) time.Time {
	// Define a gigasecond as 1 billion seconds
	gigaSecond := time.Duration(1e9) * time.Second
	
	// Add the gigasecond duration to the provided time and return the new time
	return t.Add(gigaSecond)
}
