package strand

var toReplace = map[byte]string {
    'G' : "C",
    'C' : "G",
    'T' : "A",
    'A' : "U",
}

func ToRNA(dna string) string {
    res := ""
	for i := range dna {
        rep, ok := toReplace[dna[i]]
        if ok {
            res += rep
        } else {
            res += string(dna[i])
        }
    }
    return res
}
