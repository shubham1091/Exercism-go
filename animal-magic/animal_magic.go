package chance

import (
	"math/rand"
)

// animals holds the constant list of available animals
var animals = []string{
	"ant",
	"beaver",
	"cat",
	"dog",
	"elephant",
	"fox",
	"giraffe",
	"hedgehog",
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	const (
		minDieValue = 1
		maxDieValue = 20
	)
	return rand.Intn(maxDieValue) + minDieValue
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	const maxEnergy = 12.0
	return rand.Float64() * maxEnergy
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	// Create a copy to avoid modifying the original slice
	result := make([]string, len(animals))
	copy(result, animals)

	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}
