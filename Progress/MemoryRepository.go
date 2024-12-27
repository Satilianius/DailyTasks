package Progress

import (
	"DailyTasks/Tasks"
	"errors"
	"github.com/google/uuid"
	"sync"
	"time"
)

type MemoryRepository struct {
	booleanTaskProgress map[uuid.UUID]BooleanProgress
	numberTaskProgress  map[uuid.UUID]NumberProgress
	// TODO mutex per map?
	mu sync.Mutex // For thread safety
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		booleanTaskProgress: make(map[uuid.UUID]BooleanProgress),
		numberTaskProgress:  make(map[uuid.UUID]NumberProgress),
	}
}

func (r *MemoryRepository) AddBooleanTask(task Tasks.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.booleanTaskProgress[task.Uuid] = BooleanProgress{}
	return nil
}

func (r *MemoryRepository) AddNumberTask(task Tasks.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.numberTaskProgress[task.Uuid] = NumberProgress{}
	return nil
}

func (r *MemoryRepository) GetByUuid(taskUuid uuid.UUID) (PrintableProgress, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// TODO handle missing key
	// TODO should caller know the task type? Here and in delete
	if progress, exists := r.booleanTaskProgress[taskUuid]; exists {
		return progress, nil
	}
	if progress, exists := r.numberTaskProgress[taskUuid]; exists {
		return progress, nil
	}
	return nil, errors.New("progress not found")
}

func (r *MemoryRepository) UpdateBooleanProgress(taskUuid uuid.UUID, date time.Time, done bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.booleanTaskProgress[taskUuid].DatesToValue[date] = done
	return nil
}

func (r *MemoryRepository) UpdateNumberProgress(taskUuid uuid.UUID, date time.Time, value float64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.numberTaskProgress[taskUuid].DatesToValue[date] = value
	return nil
}

func (r *MemoryRepository) Remove(taskUuid uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.booleanTaskProgress, taskUuid)
	delete(r.numberTaskProgress, taskUuid)
	return nil
}
