package reverse

func Reverse(input string) string {
	rev := ""
    runeInput := []rune(input)

    for i:=len(runeInput)-1;i>=0;i-- {
        rev += string(runeInput[i])
    }

    return rev
}
