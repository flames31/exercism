package lsproduct

import (
    "errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
    if span < 0 || span > len(digits) {
        return 0, errors.New("invalid span")
    }
	var res int64

    for i:=0 ; i+span <= len(digits) ; i++ {
        var prod int64 = 1
		for _, r := range digits[i:i+span] {
            if r < '0' || r > '9' {
                return 0, errors.New("invalid digit")
            }

            if r == '0' {
                prod = 0
                break
            }

            prod *= int64(r - '0')
        }

        if prod > res {
            res = prod
        }
    }

    return res, nil
}