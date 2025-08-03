package anagram

import (
    "sort"
    "strings"
)

func Detect(subject string, candidates []string) []string {
	res := make([]string, 0)
    
	toFind := []rune(strings.ToLower(subject))
	sort.Slice(toFind, func(i, j int) bool {
        return toFind[i] < toFind[j]
    })
    toFindStr := string(toFind)

    for _, s := range candidates {
        sRune := []rune(strings.ToLower(s))
        sort.Slice(sRune, func(i, j int) bool {
            return sRune[i] < sRune[j]
        })
        if string(sRune) == toFindStr && strings.ToLower(s) != strings.ToLower(subject) {
            res = append(res, s)
        }
    }

    return res
}
