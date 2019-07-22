package timetagger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var (
	colorGreenBold = color.New(color.FgGreen, color.Bold)
)

type timeTagger struct {
	start    time.Time
	preTag   string
	sameLine int
}

func NewTimeTagger() *timeTagger {
	return &timeTagger{start: time.Now()}
}

func (tt *timeTagger) Tag() string {
	diff := time.Now().Sub(tt.start).Round(time.Second)
	hour := diff / time.Hour

	diff -= (hour * time.Hour)
	min := diff / time.Minute

	diff -= (min * time.Minute)
	sec := diff / time.Second

	tag := colorGreenBold.Sprintf(fmt.Sprintf("%02d:%02d:%02d", hour, min, sec))

	if tag != tt.preTag {
		tt.preTag = tag
		tt.sameLine = 1
	} else {
		tt.sameLine++
		tag = color.GreenString(fmt.Sprintf("%+8d", tt.sameLine))
	}

	return tag
}
