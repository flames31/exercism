package thefarm

import (
    "errors"
    "fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, numCows int) (float64, error) {
    totalFod, err := fodderCalculator.FodderAmount(numCows)
    if err != nil {
        return 0, err
    }

    factor, err := fodderCalculator.FatteningFactor()
    if err != nil {
        return 0, err
    }

    return (totalFod*factor)/float64(numCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, numCows int) (float64, error) {
    if numCows <= 0 {
        return 0, errors.New("invalid number of cows")
    }

    food, err := DivideFood(fodderCalculator, numCows)
	if err != nil {
        return 0, err
    }

    return food, nil
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
    numCows int
    errMsg string
}

func (i *InvalidCowsError) Error() string {
    return fmt.Sprintf("%v cows are invalid: %v", i.numCows, i.errMsg)
}

func ValidateNumberOfCows(numCows int) error {
    if numCows < 0 {
        return &InvalidCowsError{
            numCows : numCows,
            errMsg : "there are no negative cows",
        }
    }

    if numCows == 0 {
        return &InvalidCowsError{
            numCows : numCows,
            errMsg : "no cows don't need food",
        }
    }

    return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
