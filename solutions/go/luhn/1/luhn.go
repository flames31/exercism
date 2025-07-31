package luhn

import "strings"

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
    if len(id) < 2 { 
    	return false
	}
	sum := 0
    
    for i:=len(id)-2; i>=0;i -= 2 {
        num, ok := numMap[id[i]]
        if !ok {
            return false
        }
        num *= 2
        if num > 9 {
            num -= 9
        }

        num2, ok := numMap[id[i+1]]
		if !ok {
            return false
        }
        sum += num + num2
    }

    if len(id) % 2 != 0 {
        sum += numMap[id[0]]
    }

    return sum % 10 == 0
}

var numMap = map[byte]int{
    '0' : 0,
    '1' : 1,
    '2' : 2,
    '3' : 3,
    '4' : 4,
    '5' : 5,
    '6' : 6,
    '7' : 7,
    '8' : 8,
    '9' : 9,
}