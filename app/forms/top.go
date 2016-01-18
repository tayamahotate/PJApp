package forms
import "time"

type Top struct {
	Id         int64
	StartDay   time.Time
	EndDay     time.Time
	Status     string
	UpdateTime time.Time

}