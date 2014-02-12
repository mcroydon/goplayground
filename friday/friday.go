// Package friday answers the age-old question "is it Friday?" It cannot assist you in which seat you should take.
package friday

import (
	"time"
)

// Friday lets you know if it is Friday for a specific time.Time.
func Friday(t time.Time) bool {
	if t.Weekday() == time.Friday {
		return true
	} else {
		return false
	}
}

// NowFriday tells you whether or not it's Friday in your local time right now.
func NowFriday() bool {
	return Friday(time.Now())
}
