package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
    letterMap := make(map[string]int)
	for k , v := range in {
        for _, s := range v {
            letterMap[strings.ToLower(s)] = k
        }
    }

    return letterMap
}
