package pangram

import (
    "strings"
)

func IsPangram(input string) bool {
	alphaSet := make(map[rune]struct{})
    count := 0

    for _, uStr := range strings.ToUpper(input) {
        if !(uStr >= 'A' && uStr <= 'Z') {
            continue
        }
		if _, ok := alphaSet[uStr]; !ok {
            count++
        }

		alphaSet[uStr] = struct{}{}        
    }

    return count == 26
}
