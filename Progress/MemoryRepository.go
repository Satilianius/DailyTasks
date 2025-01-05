package Progress

import (
	"DailyTasks/Tasks"
	"errors"
	"github.com/google/uuid"
	"sync"
	"time"
)

type MemoryRepository struct {
	booleanTaskProgress  map[uuid.UUID]*BooleanProgress
	numberTaskProgress   map[uuid.UUID]*NumberProgress
	durationTaskProgress map[uuid.UUID]*DurationProgress

	booleanMutex  sync.RWMutex
	numberMutex   sync.RWMutex
	durationMutex sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		booleanTaskProgress: make(map[uuid.UUID]*BooleanProgress),
		numberTaskProgress:  make(map[uuid.UUID]*NumberProgress),
	}
}

func (r *MemoryRepository) AddBooleanTask(task Tasks.Task) error {
	r.booleanMutex.Lock()
	defer r.booleanMutex.Unlock()

	r.booleanTaskProgress[task.Uuid] = NewBooleanProgress()
	return nil
}

func (r *MemoryRepository) AddNumberTask(task Tasks.Task) error {
	r.numberMutex.Lock()
	defer r.numberMutex.Unlock()

	r.numberTaskProgress[task.Uuid] = NewNumberProgress()
	return nil
}

func (r *MemoryRepository) AddDurationTask(task Tasks.Task) error {
	r.durationMutex.Lock()
	defer r.durationMutex.Unlock()

	r.durationTaskProgress[task.Uuid] = NewDurationProgress()
	return nil
}

func (r *MemoryRepository) GetPrintableProgressByUuid(taskUuid uuid.UUID) (PrintableProgress, error) {
	booleanProgress, found, err := r.GetBooleanByUuid(taskUuid)
	if found {
		return booleanProgress, err
	}

	numberProgress, found, err := r.GetNumberByUuid(taskUuid)
	if found {
		return numberProgress, err
	}

	return nil, errors.New("progress not found")
}

func (r *MemoryRepository) GetBooleanByUuid(taskUuid uuid.UUID) (*BooleanProgress, bool, error) {
	r.booleanMutex.RLock()
	defer r.booleanMutex.RUnlock()

	if progress, exists := r.booleanTaskProgress[taskUuid]; exists {
		return progress, true, nil
	}

	return nil, false, nil
}

func (r *MemoryRepository) GetNumberByUuid(taskUuid uuid.UUID) (*NumberProgress, bool, error) {
	r.numberMutex.RLock()
	defer r.numberMutex.RUnlock()

	if progress, exists := r.numberTaskProgress[taskUuid]; exists {
		return progress, true, nil
	}

	return nil, false, nil
}

func (r *MemoryRepository) GetDurationByUuid(taskUuid uuid.UUID) (*DurationProgress, bool, error) {
	r.durationMutex.RLock()
	defer r.durationMutex.RUnlock()

	if progress, exists := r.durationTaskProgress[taskUuid]; exists {
		return progress, true, nil
	}
	return nil, false, nil
}

func (r *MemoryRepository) UpdateBooleanProgress(taskUuid uuid.UUID, date time.Time, done bool) error {
	r.booleanMutex.Lock()
	defer r.booleanMutex.Unlock()

	if progress, exists := r.booleanTaskProgress[taskUuid]; exists {
		progress.Update(date, done)
		return nil
	}
	return errors.New("progress not found")
}

func (r *MemoryRepository) UpdateNumberProgress(taskUuid uuid.UUID, date time.Time, value float64) error {
	r.numberMutex.Lock()
	defer r.numberMutex.Unlock()

	if progress, exists := r.numberTaskProgress[taskUuid]; exists {
		progress.Update(date, value)
		return nil
	}
	return errors.New("progress not found")
}

func (r *MemoryRepository) UpdateDurationProgress(taskUuid uuid.UUID, date time.Time, value time.Duration) error {
	r.durationMutex.Lock()
	defer r.durationMutex.Unlock()

	if progress, exists := r.durationTaskProgress[taskUuid]; exists {
		progress.Update(date, value)
	}

	return errors.New("progress not found")
}

func (r *MemoryRepository) Remove(taskUuid uuid.UUID) error {
	r.RemoveBoolean(taskUuid)
	r.RemoveNumber(taskUuid)
	r.RemoveDuration(taskUuid)

	return nil
}

func (r *MemoryRepository) RemoveBoolean(taskUuid uuid.UUID) {
	r.booleanMutex.Lock()
	defer r.booleanMutex.Unlock()

	delete(r.booleanTaskProgress, taskUuid)
}

func (r *MemoryRepository) RemoveNumber(taskUuid uuid.UUID) {
	r.numberMutex.Lock()
	defer r.numberMutex.Unlock()

	delete(r.numberTaskProgress, taskUuid)
}

func (r *MemoryRepository) RemoveDuration(taskUuid uuid.UUID) {
	r.durationMutex.Lock()
	defer r.durationMutex.Unlock()

	delete(r.durationTaskProgress, taskUuid)
}
