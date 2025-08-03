package atbash

import (
    "strings"
    "unicode"
)

func Atbash(s string) string {
    s = strings.ToLower(s)
	var builder strings.Builder
    count := 0
	
    for _, c := range s {
        if !unicode.IsDigit(c) && !unicode.IsLetter(c) {
            continue
        }

        if count > 0 && count % 5 == 0 {
            builder.WriteRune(' ')
        }

		if unicode.IsLetter(c) {
            c = 'a' + ('z' - c)
        }
        
        builder.WriteRune(c)
        count++
    }

    return builder.String()
}