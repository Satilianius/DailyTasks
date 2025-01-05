package Progress

import (
	"fmt"
	"sync"
	"time"
)

type NumberProgress struct {
	datesToValue map[time.Time]float64
	mu           sync.RWMutex
}

func NewNumberProgress() *NumberProgress {
	return &NumberProgress{datesToValue: make(map[time.Time]float64)}
}

func (p *NumberProgress) Update(updateDate time.Time, updateValue float64) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.datesToValue[updateDate] = updateValue
}

func (p *NumberProgress) GetValueAtDate(day time.Time) float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.datesToValue[day]
}

func (p *NumberProgress) GetPrintableProgressAtDate(utcDate time.Time) string {
	return fmt.Sprintf("%f", p.datesToValue[utcDate])
}
