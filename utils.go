package tdiff

import (
	"strings"
	"time"

	"github.com/mjwhitta/errors"
)

func getTime(ts string) (t time.Time, e error) {
	switch ts {
	case "now":
		return time.Now(), nil
	}

	if !strings.Contains(ts, "T") {
		ts += "T00:00:00"
	}

	if t, e = time.Parse(time.RFC3339, ts+"Z"); e != nil {
		e = errors.Newf("failed to parse timestamp %s: %w", ts, e)
	}

	return
}
