// Package timetoys pants.

package timetoys

import (
	"fmt"
	"time"
)

// HowFast oops through time.Now() as fast as possible, keeping track of the maximum delta in nanoseconds and periodically
// printing statistics to stdout.
func HowFast() {
	c := 0
	delta := 0
	maxDelta := 0
	prevTime := time.Now()

	// Periodically print stats to the console.
	tick := time.NewTicker(time.Second)
	go func() {
		for t := range tick.C {
			fmt.Printf("%v processed %v current delta %v max delta %v\n", t.Format(time.StampNano), c, delta, maxDelta)
		}
	}()

	for {
		// Calculate the current time.
		now := time.Now()
		// Calculate the nanosecond difference
		newDelta := now.Nanosecond() - prevTime.Nanosecond()

		// Add the second difference if needed
		secondDelta := now.Second() - prevTime.Second()
		if secondDelta != 0 {
			newDelta = newDelta + secondDelta*1000000000
		}

		// Update statistics
		delta = newDelta
		if newDelta > maxDelta {
			maxDelta = newDelta
			fmt.Printf("%v New max delta %v since %v\n", now.Format(time.StampNano), maxDelta, prevTime.Format(time.StampNano))
		}
		prevTime = now
		c++
	}
}
