package lasagna

func PreparationTime(layers []string, averagePreparationTime int) int {
	if averagePreparationTime < 1 {
		averagePreparationTime = 2
	}

	return len(layers) * averagePreparationTime
}

func Quantities(layers []string) (int, float64) {
	totalNoodleWeightInGrams := 0
	totalSauceVolumeInLiters := 0.0

	for _, layer := range layers {
		switch layer {
		case "noodles":
			totalNoodleWeightInGrams += 50
		case "sauce":
			totalSauceVolumeInLiters += 0.2
		}
	}
	return totalNoodleWeightInGrams, totalSauceVolumeInLiters
}

func AddSecretIngredient(friendsRecipe, myRecipe []string) {
	myRecipe[len(myRecipe)-1] = friendsRecipe[len(friendsRecipe)-1]
}

func ScaleRecipe(quantities []float64, scale int) []float64 {
	baseScale := 2.0
	// Create a new slice with the same length as quantities to avoid altering the input slice.
	scaledRecipe := make([]float64, len(quantities))
	copy(scaledRecipe, quantities)

	for i := range scaledRecipe {
		scaledRecipe[i] *= float64(scale) / baseScale
	}

	return scaledRecipe
}
