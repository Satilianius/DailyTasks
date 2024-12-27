package Progress

import (
	"fmt"
	"time"
)

type BooleanProgress struct {
	DatesToValue map[time.Time]bool
}

func (b BooleanProgress) GetPrintableProgressAtDate(utcDate time.Time) string {
	if value, exists := b.DatesToValue[utcDate]; exists {
		return fmt.Sprintf("%t", value)
	}
	return "false"
}
