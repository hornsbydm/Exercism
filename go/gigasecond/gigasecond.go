package gigasecond

import "time"

var giga, _ = time.ParseDuration("1000000000s")

//AddGigasecond returns the time 1 gigasecond from the provided.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(giga)
}
