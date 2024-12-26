package Progress

import (
	"DailyTasks/Tasks"
	"time"
)

type BooleanProgress struct {
	task             Tasks.Task
	utcDatesWhenDone map[time.Time]struct{}
}

func (b BooleanProgress) GetPrintableProgressAtDate(utcDate time.Time) string {
	if _, done := b.utcDatesWhenDone[utcDate]; done {
		return "true"
	}
	return "false"
}
