package strain

// Implement the "Keep" and "Discard" function in this file.

func Keep[T any](xs []T, f func(T) bool) []T {
	if f == nil {
		return nil
	}
	
	result := make([]T, 0)
	for _, x := range xs {
		if f(x) {
			result = append(result, x)
		}
	}
	return result
}

func Discard[T any](xs []T, f func(T) bool) []T {
	if f == nil {
		return nil
	}

	result := make([]T, 0)
	for _, x := range xs {
		if !f(x) {
			result = append(result, x)
		}
	}
	return result
}

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
