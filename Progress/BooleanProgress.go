package Progress

import (
	"fmt"
	"time"
)

type BooleanProgress struct {
	datesToValue map[time.Time]bool // Missing date considered false
}

func NewBooleanProgress(datesToValue map[time.Time]bool) *BooleanProgress {
	datesToValueCopy := make(map[time.Time]bool, len(datesToValue))
	for k, v := range datesToValue {
		datesToValueCopy[k] = v
	}
	return &BooleanProgress{datesToValue: datesToValueCopy}
}

func (p *BooleanProgress) GetValueAtDate(day time.Time) bool {
	return p.datesToValue[day]
}

func (p *BooleanProgress) GetPrintableProgressAtDate(utcDate time.Time) string {
	return fmt.Sprintf("%t", p.datesToValue[utcDate])
}
