package Progress

import (
	"DailyTasks/Tasks"
	"errors"
	"github.com/google/uuid"
	"sync"
	"time"
)

type MemoryRepository struct {
	booleanTaskProgress map[uuid.UUID]*BooleanProgress
	numberTaskProgress  map[uuid.UUID]*NumberProgress

	booleanMutex sync.RWMutex // For thread safety
	numberMutex  sync.RWMutex // For thread safety
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

	r.booleanTaskProgress[task.Uuid] = &BooleanProgress{}
	return nil
}

func (r *MemoryRepository) AddNumberTask(task Tasks.Task) error {
	r.numberMutex.Lock()
	defer r.numberMutex.Unlock()

	r.numberTaskProgress[task.Uuid] = &NumberProgress{}
	return nil
}

func (r *MemoryRepository) GetPrintableProgressByUuid(taskUuid uuid.UUID) (PrintableProgress, error) {
	booleanProgress, err, found := r.GetBooleanByUuid(taskUuid)
	if found {
		return *booleanProgress, err
	}

	numberProgress, err, found := r.getNumberByUuid(taskUuid)
	if found {
		return numberProgress, err
	}

	return nil, errors.New("progress not found")
}

func (r *MemoryRepository) GetBooleanByUuid(taskUuid uuid.UUID) (*BooleanProgress, error, bool) {
	r.booleanMutex.RLock()
	defer r.booleanMutex.RUnlock()

	if progress, exists := r.booleanTaskProgress[taskUuid]; exists {
		return progress, nil, true
	}
	return nil, nil, false
}

func (r *MemoryRepository) getNumberByUuid(taskUuid uuid.UUID) (*NumberProgress, error, bool) {
	r.numberMutex.RLock()
	defer r.numberMutex.RUnlock()

	if progress, exists := r.numberTaskProgress[taskUuid]; exists {
		return progress, nil, true
	}

	return nil, nil, false
}

func (r *MemoryRepository) UpdateBooleanProgress(taskUuid uuid.UUID, date time.Time, done bool) error {
	r.booleanMutex.Lock()
	defer r.booleanMutex.Unlock()

	r.booleanTaskProgress[taskUuid].DatesToValue[date] = done
	return nil
}

func (r *MemoryRepository) UpdateNumberProgress(taskUuid uuid.UUID, date time.Time, value float64) error {
	r.numberMutex.Lock()
	defer r.numberMutex.Unlock()

	r.numberTaskProgress[taskUuid].DatesToValue[date] = value
	return nil
}

func (r *MemoryRepository) Remove(taskUuid uuid.UUID) error {
	r.RemoveBoolean(taskUuid)
	r.RemoveNumber(taskUuid)

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
