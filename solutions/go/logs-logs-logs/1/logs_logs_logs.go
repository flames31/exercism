package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	charMap := map[int32]string{'❗':"recommendation",'🔍':"search",'☀':"weather"}

    for _, val := range log{
        if app, exists := charMap[val]; exists == true {
            return app
        }
    }

    return "default"
    
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newString := ""
    for _, val := range log {
        if val == oldRune {
            val = newRune
        }

        newString += string(val)
    }

    return newString
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	numberOfRunes := utf8.RuneCountInString(log)

    if numberOfRunes <= limit {
        return true
    } else {
        return false
    }
}
