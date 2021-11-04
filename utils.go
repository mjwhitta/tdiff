package tdiff

import (
	"strings"
	"time"
)

func getTime(ts string) (time.Time, error) {
	switch ts {
	case "now":
		return time.Now(), nil
	}

	if !strings.Contains(ts, "T") {
		ts += "T00:00:00"
	}

	return time.Parse(time.RFC3339, ts+"Z")
}
