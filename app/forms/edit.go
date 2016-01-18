package forms

import "time"

type Edit struct {
	Id         uint64
	StartDay   time.Time
	EndDay     time.Time
	Status     string
	UpdateTime time.Time

}