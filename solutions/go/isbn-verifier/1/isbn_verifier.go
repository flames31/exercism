package isbn

import (
    "strings"
    "strconv"
)

func IsValidISBN(isbn string) bool {
    isbn = strings.ReplaceAll(isbn, "-", "")

    if len(isbn) != 10 {
        return false
    }
    i := 10
    count := 0
	for idx, c := range isbn {
        var val int
        if c == 'X' {
            if idx != len(isbn)-1 {
                return false
            }
            val = 10
        } else {
            ele, err := strconv.Atoi(string(c))
            if err != nil {
                return false
            }

            val = ele
        }

        count += (val * i)
        i--
    }

    return count % 11 == 0
}
