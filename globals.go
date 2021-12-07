package tdiff

import (
	"regexp"
	"time"
)

// Day is a Duration of 24 hours
const Day time.Duration = 24 * time.Hour

// RFC3339 is a regex to validate user-input.
var RFC3339 *regexp.Regexp = regexp.MustCompile(
	`^(\d{4}-\d\d-\d\d(T\d\d:\d\d:\d\d)?|now)$`,
)

// Version is the package version
const Version = "1.0.1"
