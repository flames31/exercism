package raindrops

import "strconv"

func Convert(number int) string {
	if number % 3 != 0 && number % 5 != 0 && number % 7 != 0 {
        return strconv.Itoa(number)
    }

    res := ""

    if number % 3 == 0 { res += "Pling" }
    if number % 5 == 0 { res += "Plang" }
    if number % 7 == 0 { res += "Plong" }

    return res
}
