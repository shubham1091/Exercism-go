package thefarm

import "fmt"

type InvalidCowsError struct {
	Cows int
	Msg  string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.Cows, e.Msg)
}

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	fodderAmount, err := fc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return fodderAmount / float64(cows) * fatteningFactor, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, fmt.Errorf("invalid number of cows")
	}
	return DivideFood(fc, cows)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	switch {
	case cows < 0:
		return &InvalidCowsError{
			Cows: cows,
			Msg:  "there are no negative cows",
		}
	case cows == 0:
		return &InvalidCowsError{
			Cows: cows,
			Msg:  "no cows don't need food",
		}
	default:
		return nil
	}
}
