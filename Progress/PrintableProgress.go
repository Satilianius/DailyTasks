package Progress

import "time"

type PrintableProgress interface {
	GetPrintableProgressAtDate(time time.Time) string
}
