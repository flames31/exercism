package clock

import (
    "time"
    "fmt"
)

// Define the Clock type here.
type Clock struct {
    t time.Time
}

func New(h, m int) Clock {
    return Clock{
        t : normalize(h, m),
    }
}

func (c Clock) Add(m int) Clock {
	return Clock{
        t : normalize(c.t.Hour(), c.t.Minute() + m),
    }
}

func (c Clock) Subtract(m int) Clock {
	return Clock{
        t : normalize(c.t.Hour(), c.t.Minute() - m),
    }
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.t.Hour(), c.t.Minute())
}

func normalize(h, m int) time.Time {
    totalMinutes := ((h*60 + m) % (24 * 60) + (24 * 60)) % (24 * 60)
	hour := totalMinutes / 60
	min := totalMinutes % 60
    return time.Date(0, 1, 1, hour, min, 0, 0, time.UTC) 
}
