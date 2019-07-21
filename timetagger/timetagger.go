package timetagger

import (
	"fmt"
	"time"
)

type TimeTagger struct {
	start time.Time
	val   string
}

func NewTimeTagger() *TimeTagger {
	return &TimeTagger{start: time.Now()}
}

func (tt *TimeTagger) Tag() string {
	diff := time.Now().Sub(tt.start).Round(time.Second)
	hour := diff / time.Hour

	diff -= (hour * time.Hour)
	min := diff / time.Minute

	diff -= (min * time.Minute)
	sec := diff / time.Second

	val := fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)

	if val != tt.val {
		tt.val = val
	} else {
		val = ""
	}

	return val
}
