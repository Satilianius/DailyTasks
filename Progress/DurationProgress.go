package Progress

import (
	"fmt"
	"sync"
	"time"
)

type DurationProgress struct {
	datesToValue map[time.Time]time.Duration
	mu           sync.RWMutex
}

func NewDurationProgress() *DurationProgress {
	return &DurationProgress{datesToValue: make(map[time.Time]time.Duration)}
}

func (p *DurationProgress) Update(updateDate time.Time, updateValue time.Duration) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.datesToValue[updateDate] = updateValue
}

func (p *DurationProgress) GetValueAtDate(day time.Time) time.Duration {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.datesToValue[day]
}

func (p *DurationProgress) GetPrintableProgressAtDate(utcDate time.Time) string {
	return fmt.Sprintf("%s", p.datesToValue[utcDate])
}
