// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	prov := make([]string, 0)

    for i:=1;i<len(rhyme);i++ {
        prov = append(prov, fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i-1],rhyme[i]))
    }
    if len(rhyme) > 0 {
        prov = append(prov, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))
    }
    
    return prov

    
}
