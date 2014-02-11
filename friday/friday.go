// Package friday answers the age-old question "is it Friday?" It cannot assist you in which seat you should take.
package friday

import (
	"time"
)

// Friday tells you whether or not it's Friday in your local time.
func Friday() bool {
	if time.Now().Weekday() == time.Friday {
		return true
	} else {
		return false
	}
}
