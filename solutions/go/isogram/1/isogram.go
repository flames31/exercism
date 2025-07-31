package isogram

import "strings"

func IsIsogram(word string) bool {
	letterMap := make(map[rune]struct{})

    for _, w := range strings.ToLower(word) {
        if w == ' ' || w == '-' {
            continue
        }

        if _, ok := letterMap[w]; ok {
            return false
        }

        letterMap[w] = struct{}{}
    }

    return true
}
