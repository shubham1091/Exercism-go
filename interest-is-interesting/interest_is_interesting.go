package interest

const (
	// Interest rate thresholds and values
	negativeRate float32 = 3.213
	lowRate      float32 = 0.5
	mediumRate   float32 = 1.621
	highRate     float32 = 2.475

	// Balance thresholds
	lowThreshold    float64 = 1000
	mediumThreshold float64 = 5000
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return negativeRate
	case balance < lowThreshold:
		return lowRate
	case balance < mediumThreshold:
		return mediumRate
	default:
		return highRate
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	rateOfInterest := InterestRate(balance)
	return balance * float64(rateOfInterest) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	if balance >= targetBalance {
		return 0
	}

	years := 0
	currentBalance := balance

	for currentBalance < targetBalance {
		currentBalance = AnnualBalanceUpdate(currentBalance)
		years++

		// Prevent infinite loop for cases where interest is negative
		if years > 1000 {
			return 1000
		}
	}

	return years
}
