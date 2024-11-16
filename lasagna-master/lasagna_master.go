package lasagna

// Constants for ingredient quantities per layer
const (
	noodlesGram = 50  // Amount of noodles in grams per layer
	sauceLiters = 0.2 // Amount of sauce in liters per layer
)

// PreparationTime calculates the total preparation time based on the number of layers
// and the time required per layer. If prepTime is not provided (0), a default of 2 minutes per layer is used.
func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2 // Default preparation time per layer
	}
	return len(layers) * prepTime
}

// Quantities calculates the total quantity of noodles (in grams) and sauce (in liters)
// needed based on the layers provided. Each "noodles" layer adds 50g, and each "sauce" layer adds 0.2L.
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += noodlesGram
		case "sauce":
			sauce += sauceLiters
		}
	}
	return noodles, sauce
}

// AddSecretIngredient replaces the last ingredient in myList with the last ingredient
// from friendsList. This assumes myList already has a placeholder for the secret ingredient.
func AddSecretIngredient(friendsList, myList []string) {
	// Replace the last element of myList with the last element of friendsList
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe scales the quantities of ingredients based on the desired number of portions.
// The original recipe is for 2 portions. Returns a new slice with scaled quantities.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities)) // Create a new slice for the scaled quantities
	scaleFactor := float64(portions) / 2.0               // Scaling factor relative to the base 2 portions

	for i, quantity := range quantities {
		scaledQuantities[i] = quantity * scaleFactor
	}
	return scaledQuantities
}
