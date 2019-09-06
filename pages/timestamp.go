package pages

import (
	"fmt"
	"time"
)

type Timestamp int64

func (ts Timestamp) Since() string {
	// timestamp is in s
	amt, unit := func() (float64, string) {
		posttime := time.Unix(int64(ts), 0)
		since := time.Since(posttime)
		hrs := since.Hours()
		switch {
		case hrs < 1:
			if mins := since.Minutes(); mins < 1 {
				return since.Seconds(), "second"
			} else {
				return mins, "minute"
			}
		case hrs < 24:
			return hrs, "hour"
		default:
			return hrs / 24, "day"
		}
	}()
	if int(amt) != 1 {
		unit += "s"
	}
	return fmt.Sprintf(" %v %v ago", int(amt), unit)
}
