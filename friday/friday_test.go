package friday

import (
	"testing"
	"time"
)

func TestFriday(t *testing.T) {
	utc := time.UTC
	fridays := []time.Time{time.Date(2014, 2, 14, 12, 0, 0, 0, utc),
		time.Date(2013, 11, 22, 5, 0, 0, 0, utc),
		time.Date(1985, 10, 18, 10, 0, 0, 0, utc)}
	nonFridays := []time.Time{time.Date(2014, 1, 1, 0, 0, 0, 0, utc),
		time.Date(1991, 9, 23, 10, 0, 0, 0, utc),
		time.Date(1970, 1, 1, 0, 0, 0, 0, utc)}

	for _, ti := range fridays {
		isFriday := Friday(ti)
		if !isFriday {
			t.Errorf("Expected Friday(%s) to be true.", ti)
		}
	}

	for _, ti := range nonFridays {
		isFriday := Friday(ti)
		if isFriday {
			t.Errorf("Expected Friday(%s) to be false.", ti)
		}
	}
}

func TestNowFriday(t *testing.T) {
	weekday := time.Now().Weekday()
	isFriday := NowFriday()
	if weekday == time.Friday && !isFriday {
		t.Errorf("It appears to be Friday local time and Friday() returned false.")
	} else {
		if isFriday {
			t.Errorf("It is %v and Friday() returned true.", weekday)
		}
	}
}
