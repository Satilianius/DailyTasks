package Progress

import (
	"DailyTasks/Tasks"
	"errors"
	"github.com/google/uuid"
	"sync"
	"time"
)

type ConcurrentMemoryRepository struct {
	booleanTaskProgress sync.Map
	numberTaskProgress  sync.Map
}

func NewConcurrentMemoryRepository() *ConcurrentMemoryRepository {
	return &ConcurrentMemoryRepository{
		booleanTaskProgress: sync.Map{},
		numberTaskProgress:  sync.Map{},
	}
}

func (r *ConcurrentMemoryRepository) AddBooleanTask(task Tasks.Task) error {
	r.booleanTaskProgress.Store(task.Uuid, NewBooleanProgress())
	return nil
}

func (r *ConcurrentMemoryRepository) AddNumberTask(task Tasks.Task) error {
	r.numberTaskProgress.Store(task.Uuid, NewNumberProgress())
	return nil
}

func (r *ConcurrentMemoryRepository) GetPrintableProgressByUuid(taskUuid uuid.UUID) (PrintableProgress, error) {
	booleanProgress, found, err := r.GetBooleanByUuid(taskUuid)
	if found {
		return booleanProgress, err
	}

	numberProgress, found, err := r.getNumberByUuid(taskUuid)
	if found {
		return numberProgress, err
	}

	return nil, errors.New("progress not found")
}

func (r *ConcurrentMemoryRepository) GetBooleanByUuid(taskUuid uuid.UUID) (*BooleanProgress, bool, error) {
	if progress, exists := r.booleanTaskProgress.Load(taskUuid); exists {
		return progress.(*BooleanProgress), true, nil
	}

	return nil, false, nil
}

func (r *ConcurrentMemoryRepository) getNumberByUuid(taskUuid uuid.UUID) (*NumberProgress, bool, error) {
	if progress, exists := r.numberTaskProgress.Load(taskUuid); exists {
		return progress.(*NumberProgress), true, nil
	}

	return nil, false, nil
}

func (r *ConcurrentMemoryRepository) UpdateBooleanProgress(taskUuid uuid.UUID, date time.Time, done bool) error {
	// TODO lock?
	booleanProgress, exist := r.booleanTaskProgress.Load(taskUuid)
	if !exist {
		return errors.New("progress not found")
	}

	booleanProgress.(*BooleanProgress).DatesToValue[date] = done
	return nil
}

func (r *ConcurrentMemoryRepository) UpdateNumberProgress(taskUuid uuid.UUID, date time.Time, value float64) error {
	// TODO lock?
	numberProgress, exist := r.numberTaskProgress.Load(taskUuid)
	if !exist {
		return errors.New("progress not found")
	}
	numberProgress.(*NumberProgress).DatesToValue[date] = value
	return nil
}

func (r *ConcurrentMemoryRepository) Remove(taskUuid uuid.UUID) error {
	r.booleanTaskProgress.Delete(taskUuid)
	r.numberTaskProgress.Delete(taskUuid)

	return nil
}
