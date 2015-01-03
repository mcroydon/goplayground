package convert

import (
	"fmt"
)

func Example_usage() {
	// Convert 1 statute mile to nautical miles
	nautical := StatuteToNauticalMiles(1)
	fmt.Printf("%.2f\n", nautical)

	// Convert 1 nautical mile to statute miles
	statute := NauticalToStatuteMiles(1)
	fmt.Printf("%.2f\n", statute)

	// Output:
	// 1.15
	// 0.87
}
