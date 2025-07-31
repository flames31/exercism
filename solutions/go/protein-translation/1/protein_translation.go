package protein

import "fmt"

var codonMap = map[string]string {
    "AUG" : "Methionine",
	"UUU" :	"Phenylalanine",
    "UUC" :	"Phenylalanine",
	"UUA" : "Leucine",
    "UUG" :	"Leucine",
	"UCU" :	"Serine",
    "UCC" :	"Serine",
    "UCA" :	"Serine",
    "UCG" :	"Serine",
	"UAU" : "Tyrosine",
    "UAC" :	"Tyrosine",
	"UGU" :	"Cysteine",
    "UGC" :	"Cysteine",
	"UGG" : "Tryptophan",
	"UAA" :	"STOP",
    "UAG" :	"STOP",
    "UGA" :	"STOP",
}

var ErrStop = fmt.Errorf("stop codon")
var ErrInvalidBase = fmt.Errorf("invalid base")

func FromRNA(rna string) ([]string, error) {
    seq := make([]string, 0)
	l, r := 0, 2

    for r < len(rna) {
        val, ok := codonMap[rna[l:r+1]]
        if !ok {
            return []string{}, ErrInvalidBase
        }
    
        if val == "STOP" {
            break
        }

        seq = append(seq, val)
        l += 3
        r += 3
    }

    return seq, nil
}

func FromCodon(codon string) (string, error) {
	protein, ok := codonMap[codon]
    if !ok {
        return "", ErrInvalidBase
    }

    if protein == "STOP" {
        return "", ErrStop
    }

    return protein, nil
}
