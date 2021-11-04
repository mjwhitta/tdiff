package tdiff

import (
	"time"

	hl "gitlab.com/mjwhitta/hilighter"
)

// TDiff is a struct to hold the different between two dates for
// future use.
type TDiff struct {
	diff time.Duration
}

// New will return a new TDiff instance.
func New(ts1 string, ts2 string) (*TDiff, error) {
	var e error
	var t1 time.Time
	var t2 time.Time
	var td *TDiff

	if t1, e = getTime(ts1); e != nil {
		return nil, e
	}

	if t2, e = getTime(ts2); e != nil {
		return nil, e
	}

	td = &TDiff{diff: t2.Sub(t1)}
	if t1.After(t2) {
		td.diff = t1.Sub(t2)
	}

	return td, e
}

// Days will return a string showing the full time duration in days.
func (t *TDiff) Days() string {
	return hl.Sprintf("%0.2f days", t.diff.Hours()/24)
}

// Hours will return a string showing the full time duration in hours.
func (t *TDiff) Hours() string {
	return hl.Sprintf("%0.2f hours", t.diff.Hours())
}

// Microseconds will return a string showing the full time duration in
// microseconds.
func (t *TDiff) Microseconds() string {
	return hl.Sprintf("%d microseconds", t.diff.Microseconds())
}

// Milliseconds will return a string showing the full time duration in
// milliseconds.
func (t *TDiff) Milliseconds() string {
	return hl.Sprintf("%d milliseconds", t.diff.Milliseconds())
}

// Minutes will return a string showing the full time duration in
// minutes.
func (t *TDiff) Minutes() string {
	return hl.Sprintf("%0.2f minutes", t.diff.Minutes())
}

// Months will return a string showing the full time duration in
// months.
func (t *TDiff) Months() string {
	return hl.Sprintf("%0.2f months", t.diff.Hours()/24/(365/12))
}

// Nanoseconds will return a string showing the full time duration in
// nanoseconds.
func (t *TDiff) Nanoseconds() string {
	return hl.Sprintf("%d nanoseconds", t.diff.Nanoseconds())
}

// Seconds will return a string showing the full time duration in
// seconds.
func (t *TDiff) Seconds() string {
	return hl.Sprintf("%0.2f seconds", t.diff.Seconds())
}

// Years will return a string showing the full time duration in years.
func (t *TDiff) Years() string {
	return hl.Sprintf("%0.2f years", t.diff.Hours()/24/365)
}

// YearsShorthand will return a string showing if < 1 year or > 1
// year.
func (t *TDiff) YearsShorthand() string {
	var out = "Less than one year"

	if t.diff > time.Duration(time.Hour)*24*365 {
		out = "One or more years"
	}

	return out
}
