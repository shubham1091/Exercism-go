package expenses

import (
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	out := []Record{}
	for _, record := range in {
		if predicate(record) {
			out = append(out, record)
		}
	}
	return out
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(rc Record) bool {
		return rc.Day >= p.From && rc.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var total float64
	fc := ByDaysPeriod(p)
	for _, record := range in {
		if fc(record) {
			total += record.Amount
		}
	}
	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	catExist := false

	categoryFilter := ByCategory(c)
	peroidFilter := ByDaysPeriod(p)

	filtered := Filter(in, func(r Record) bool {
		if r.Category == c {
			catExist = true
		}
		return categoryFilter(r) && peroidFilter(r)
	})

	if !catExist {
		return 0, fmt.Errorf("no records for category %s", c)
	}
	return TotalByPeriod(filtered, p), nil
}
