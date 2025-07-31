package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
    if avgPrepTime == 0 {
        avgPrepTime = 2
    }
    return len(layers) * avgPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodlesCount := 0
    sauceCount := 0.0

    for _, val := range(layers){
        if val == "noodles"{
            noodlesCount += 50
        }
        if val == "sauce"{
            sauceCount += 0.2
        }
    }

    return noodlesCount, sauceCount
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendRecipe []string, myRecipe []string) {
    myRecipe[len(myRecipe)-1] = friendRecipe[len(friendRecipe)-1]
    return
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
    var newAmounts []float64
    for _,val := range(amounts){
        newAmounts = append(newAmounts, val*(float64(portions)/2))
    }

    return newAmounts
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
