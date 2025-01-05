package Progress

import (
	"fmt"
	"sync"
	"time"
)

type BooleanProgress struct {
	// TODO Sorted Set?
	datesToValue map[time.Time]bool // Missing date considered false
	mu           sync.RWMutex
}

func NewBooleanProgress() *BooleanProgress {
	return &BooleanProgress{datesToValue: make(map[time.Time]bool)}
}

func (p *BooleanProgress) Update(updateDate time.Time, updateValue bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.datesToValue[updateDate] = updateValue
}

func (p *BooleanProgress) GetValueAtDate(day time.Time) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.datesToValue[day]
}

func (p *BooleanProgress) GetPrintableProgressAtDate(utcDate time.Time) string {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return fmt.Sprintf("%t", p.datesToValue[utcDate])
}
