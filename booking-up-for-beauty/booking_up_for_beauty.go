package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	// Define multiple formats for parsing date
	formats := []string{
		"1/2/2006 15:04:05",
		"January 2, 2006 15:04:05",
		"Monday, January 2, 2006 15:04:05",
	}
	for _, dateFormat := range formats {
		theTime, err := time.Parse(dateFormat, date)
		if err == nil {
			return theTime
		}
	}
	// Graceful handling of unsupported format
	fmt.Printf("Error: '%s' does not match any supported date format.\n", date)
	return time.Time{}
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	return Schedule(date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	hour := Schedule(date).Hour()
	// Check if the hour is between 12 and 17 inclusive
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	// Use time.Format for consistent formatting
	return Schedule(date).Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	// Return anniversary in UTC to avoid issues with local time
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
