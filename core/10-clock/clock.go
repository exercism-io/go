// Package clock includes a solution for the "Clock" problem in the Go track on https://exercism.io.
package clock

import "fmt"

// Clock handles times without dates.
type Clock int

// MinutesInAnHour is the number of minutes in an hour
const MinutesInAnHour = 60

// HoursInADay is the number of hours in a day
const HoursInADay = 24

// MinutesInADay is the number of minutes in a day
const MinutesInADay = MinutesInAnHour * HoursInADay

// New returns a new Clock type object.
func New(hours, minutes int) Clock {
	return Clock(((hours*MinutesInAnHour+minutes)%MinutesInADay + MinutesInADay) % MinutesInADay)
}

// Add adds minutes to a clock.
func (clock Clock) Add(minutes int) Clock {
	return New(0, int(clock)+minutes)
}

// Subtract subtracts minutes from a clock.
func (clock Clock) Subtract(minutes int) Clock {
	return clock.Add(-minutes)
}

// String is the string representation of a clock.
func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", int(clock)/60, int(clock)%60)
}
