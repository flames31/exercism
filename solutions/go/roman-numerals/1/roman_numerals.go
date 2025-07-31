package romannumerals

import "errors"

func ToRomanNumeral(num int) (string, error) {
    if num <= 0 || num > 3999 {
        return "", errors.New("input out of range")
    }
	valueSlice := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    keySlice := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    finalStr := ""

    // check till number becomes 0
    for num > 0 {
        // check valueSlice in given order with number
        for idx, val := range valueSlice {
            // if number is greater than current element in valueSlice, reduce that value from number and append that character in final-string. and break from inner loop to check number with valueSlice again from 1000
            if num >= val {
                finalStr = finalStr + keySlice[idx]
                num = num - val
                break
            }
        }
    }

    return finalStr, nil
}
