package Progress

import (
	"DailyTasks/Tasks"
	"fmt"
	"time"
)

type NumberProgress struct {
	task         Tasks.Task
	DatesToValue map[time.Time]float64
}

func (n NumberProgress) GetPrintableProgressAtDate(utcDate time.Time) string {
	return fmt.Sprintf("%f", n.DatesToValue[utcDate])
}
