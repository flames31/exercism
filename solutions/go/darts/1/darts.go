package darts

func Score(x, y float64) int {
	d := x*x + y*y

    if d <= 1 {
        return 10
    }

    if d <= 25 {
        return 5
    }

    if d <= 100 {
        return 1
    }

    return 0
}
