package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
    if n <= 0 {
        return 0, errors.New("invalid n")
    }
	candidate := 2
	primes := make([]int, 0)

    for len(primes) < n {
        if isPrime(candidate, primes) {
            primes = append(primes, candidate)
        }

        candidate++
    }

    return primes[n-1], nil
}

func isPrime(n int, primes []int) bool {
    for _, p := range primes {
        if p > n {
            break
        }

        if n % p == 0 {
            return false
        }
    }

    return true
}