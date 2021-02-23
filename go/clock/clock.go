// Package clock implements time.
package clock

import "fmt"

const (
	minutesInHour = 60
	minutesInDay  = 1440
)

// Clock represents a 24 hour clock.
type Clock struct {
	m int //Minutes; store in simplest unit of measure.
}

// New creates a Clock.
func New(h int, m int) (c Clock) {
	c = Clock{(h*minutesInHour + m) % minutesInDay}
	c.normalize()
	return c
}

// Subtract decrements clock by m minutes.
func (c Clock) Subtract(m int) Clock {
	c.m -= m
	c.normalize()
	return c
}

// Add increments clock by m minutes.
func (c Clock) Add(m int) Clock {
	c.m += m
	c.normalize()
	return c
}

// String represents time as string in the format HH:MM.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.m/minutesInHour, c.m%minutesInHour)
}

// Normalize Clock time to 24 hour time.
func (c *Clock) normalize() {
	c.m %= 1440
	if c.m < 0 {
		c.m += minutesInDay
	}
	return
}
