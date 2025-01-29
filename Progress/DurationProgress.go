package Progress

import (
	"fmt"
	"time"
)

type DurationProgress struct {
	datesToValue map[time.Time]time.Duration
}

func NewDurationProgress(datesToValue map[time.Time]time.Duration) *DurationProgress {
	datesToValueCopy := make(map[time.Time]time.Duration, len(datesToValue))
	for k, v := range datesToValue {
		datesToValueCopy[k] = v
	}
	return &DurationProgress{datesToValue: datesToValueCopy}
}

func (p *DurationProgress) GetValueAtDate(day time.Time) time.Duration {
	return p.datesToValue[day]
}

func (p *DurationProgress) GetPrintableProgressAtDate(utcDate time.Time) string {
	return fmt.Sprintf("%s", p.datesToValue[utcDate])
}
