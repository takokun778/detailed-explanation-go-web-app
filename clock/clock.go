package clock

import "time"

type Clocker interface {
	Now() time.Time
}

type RealClocker struct{}

func (RealClocker) Now() time.Time {
	return time.Now()
}

type FixedClocker struct{}

func (FixedClocker) Now() time.Time {
	return time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
}
