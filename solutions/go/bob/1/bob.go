// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
    "strings"
    "unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
    remark = removeAllWhitespace(remark)
    if len(remark) == 0 {
        return "Fine. Be that way!"
    }

    yellCheck := isYell(remark)

    if yellCheck && remark[len(remark)-1] == '?' {
        return "Calm down, I know what I'm doing!"
    }

    if yellCheck {
        return "Whoa, chill out!"
    }

    if remark[len(remark)-1] == '?' {
        return "Sure."
    }
	return "Whatever."
}

func isYell(str string) bool {
    b := false
    for _, l := range str {
        if !unicode.IsLetter(l) {
            continue
        }
        if (l >= 'A' && l <= 'Z') {
            b = true
        } else {
            return false
        }
    }
    
    return b
}

func removeAllWhitespace(s string) string {
    return strings.Map(func(r rune) rune {
        if unicode.IsSpace(r) {
            return -1
        }
        return r
    }, s)
}