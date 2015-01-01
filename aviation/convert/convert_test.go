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

}
