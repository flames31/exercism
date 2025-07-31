package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	return colorMap[colors[0]] * 10 + colorMap[colors[1]]
}

var colorMap = map[string]int {
    "black" : 0,
    "brown": 1,
	"red": 2,
	"orange": 3,
	"yellow": 4,
	"green": 5,
	"blue": 6,
	"violet": 7,
	"grey": 8,
	"white": 9,
}