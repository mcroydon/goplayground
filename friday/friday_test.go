package friday

import (
	"testing"
	"time"
)

func TestReadJson(t *testing.T) {
	weekday := time.Now().Weekday()
	isFriday := Friday()
	if weekday == time.Friday && !isFriday {
		t.Errorf("It appears to be Friday local time and Friday() returned false.")
	} else {
		if isFriday {
			t.Errorf("It is %v and Friday() returned true.", weekday)
		}
	}
}
