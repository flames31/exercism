package wordcount

import (
    "strings"
    "unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := Frequency{}
    brokenPhrase := strings.FieldsFunc(strings.ToLower(phrase), func(c rune) bool {
        return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '\''
    })

    for _, w := range brokenPhrase {
        if w == "'" || len(w) == 0 {
            continue
        }
        w = strings.Trim(w, "'")
        freq[w]++
    }

    return freq
}
