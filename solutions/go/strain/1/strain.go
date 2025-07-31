package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](elements []T, predicate func(T) bool ) []T {
    keep := make([]T, 0)
    for _, ele := range elements {
        if predicate(ele) {
            keep = append(keep, ele)
        }
    }

    return keep
}

func Discard[T any](elements []T, predicate func(T) bool ) []T {
    discard := make([]T, 0)
    for _, ele := range elements {
        if !predicate(ele) {
            discard = append(discard, ele)
        }
    }

    return discard
}