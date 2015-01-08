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
		t.Errorf("Expected NauticalToStatuteMiles(1) to be 0.87.", sm)
	}

	lbs := fmt.Sprintf("%.2f", GallonsOfGasToPounds(1))
	if lbs != "6.00" {
		t.Errorf("Expected GallonsOfGasToPounds(1) to be 6.00.", lbs)
	}

	twolbs := fmt.Sprintf("%.2f", GallonsOfGasToPounds(2))
	if twolbs != "12.00" {
		t.Errorf("Expected GallonsOfGasToPounds(1) to be 12.00.", lbs)
	}

	gallons := fmt.Sprintf("%.2f", PoundsOfGasToGallons(12))
	if gallons != "2.00" {
		t.Errorf("Expected PoundsOfGasToGallons(1) to be 2.00.", gallons)
	}

	lbsJetA := fmt.Sprintf("%.2f", GallonsOfJetAToPounds(1))
	if lbsJetA != "6.80" {
		t.Errorf("Expected GallonsOfJetAToPounds(1) to be 6.80.", lbsJetA)
	}

	gallonsJetA := fmt.Sprintf("%.2f", PoundsOfJetAToGallons(12))
	if gallonsJetA != "1.76" {
		t.Errorf("Expected PoundsOfJetAToGallons(1) to be 1.76.", gallonsJetA)
	}

	lbsWater := fmt.Sprintf("%.2f", GallonsOfWaterToPounds(1))
	if lbsWater != "8.35" {
		t.Errorf("Expected GallonsOfWaterToPounds(1) to be 8.35.", lbsWater)
	}

	gallonsWater := fmt.Sprintf("%.2f", PoundsOfWaterToGallons(12))
	if gallonsWater != "1.44" {
		t.Errorf("Expected PoundsOfWaterToGallons(1) to be 1.44.", gallonsWater)
	}

	lbsOil := fmt.Sprintf("%.2f", GallonsOfOilToPounds(1))
	if lbsOil != "7.50" {
		t.Errorf("Expected GallonsOfOilToPounds(1) to be 7.50.", lbsOil)
	}

	gallonsOil := fmt.Sprintf("%.2f", PoundsOfOilToGallons(12))
	if gallonsOil != "1.60" {
		t.Errorf("Expected PoundsOfOilToGallons(1) to be 1.60.", gallonsOil)
	}

}
