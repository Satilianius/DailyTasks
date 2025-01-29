package Progress

import (
	"fmt"
	"time"
)

type NumberProgress struct {
	datesToValue map[time.Time]float64
}

func NewNumberProgress(datesToValue map[time.Time]float64) *NumberProgress {
	datesToValueCopy := make(map[time.Time]float64, len(datesToValue))
	for k, v := range datesToValue {
		datesToValueCopy[k] = v
	}
	return &NumberProgress{datesToValue: datesToValueCopy}
}

func (p *NumberProgress) GetValueAtDate(day time.Time) float64 {
	return p.datesToValue[day]
}

func (p *NumberProgress) GetPrintableProgressAtDate(utcDate time.Time) string {
	return fmt.Sprintf("%f", p.datesToValue[utcDate])
}
