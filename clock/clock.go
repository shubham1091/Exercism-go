package clock

import "fmt"

type Clock struct {
	h, m int
}

func New(h, m int) Clock {
	// Convert total time to minutes
	totalMinutes := h*60 + m

	// Handle negative times
	for totalMinutes < 0 {
		totalMinutes += 24 * 60
	}

	// Normalize to 24-hour format
	totalMinutes = totalMinutes % (24 * 60)

	// Convert back to hours and minutes
	hours := (totalMinutes / 60) % 24
	minutes := totalMinutes % 60

	return Clock{hours, minutes}
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
