package sorting

import (
    "fmt"
    "strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1d.0", nb.Number())
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
    res := 0
    if _, ok := fnb.(FancyNumber); ok {
        res, _ = strconv.Atoi(fnb.Value())
    }
    return res
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
    var val int 
	switch x := fnb.(type) {
        case FancyNumber : val = ExtractFancyNumber(x)
        default : val = 0
    }

    return fmt.Sprintf("This is a fancy box containing the number %v.0", val)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch x := i.(type) {
        case int : return DescribeNumber(float64(x))
        case float64 : return DescribeNumber(x)
        case NumberBox : return DescribeNumberBox(x)
        case FancyNumberBox : return DescribeFancyNumberBox(x)
        default : return "Return to sender"
    }
}
