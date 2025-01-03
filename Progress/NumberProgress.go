package Progress

import (
	"fmt"
	"time"
)

type NumberProgress struct {
	DatesToValue map[time.Time]float64
}

func NewNumberProgress() *NumberProgress {
	return &NumberProgress{DatesToValue: make(map[time.Time]float64)}
}

func (n NumberProgress) GetPrintableProgressAtDate(utcDate time.Time) string {
	return fmt.Sprintf("%f", n.DatesToValue[utcDate])
}
