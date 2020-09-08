package utils

import "time"

// UnixDuration is a wrapper around time.Unix that accepts
// a duration.
func UnixDuration(d time.Duration) time.Time {
	return time.Unix(0, 0).Add(d)
}
