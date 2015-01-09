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

	// Convert 1 gallon of gas (100LL) to pounds.
	lbs := GallonsOfGasToPounds(1)
	fmt.Printf("%.2f\n", lbs)

	// Output:
	// 1.15
	// 0.87
	// 6.00
}
