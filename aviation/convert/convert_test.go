package convert

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {

	nm := fmt.Sprintf("%.2f", StatuteToNauticalMiles(1))
	if nm != "1.15" {
		t.Errorf("Expected StatuteToNauticalMiles(1) to be 1.15.", nm)
	}

	sm := fmt.Sprintf("%.2f", NauticalToStatuteMiles(1))
	if sm != "0.87" {
		t.Errorf("Expected NauticalToStatuteMiles(1) to be 0.87.", nm)
	}

	lbs := fmt.Sprintf("%.2f", GallonsOfGasToPounds(1))
	if lbs != "6.00" {
		t.Errorf("Expected GallonsOfGasToPounds(1) to be 6.00.", nm)
	}

	gallons := fmt.Sprintf("%.2f", PoundsOfGasToGallons(12))
	if gallons != "2.00" {
		t.Errorf("Expected PoundsOfGasToGallons(1) to be 2.00.", nm)
	}

}
