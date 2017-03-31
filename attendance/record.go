package attendance

import "time"

type Record struct {
	time.Time
	LocalAuth Passer
	Attendee  Attendee
}
