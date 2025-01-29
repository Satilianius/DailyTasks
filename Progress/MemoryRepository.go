package Progress

import (
	"DailyTasks/Tasks"
	"fmt"
	"github.com/google/uuid"
	"sync"
	"time"
)

type MemoryRepository struct {
	booleanTaskProgress  map[uuid.UUID]map[time.Time]bool // Missing date considered false
	numberTaskProgress   map[uuid.UUID]map[time.Time]float64
	durationTaskProgress map[uuid.UUID]map[time.Time]time.Duration

	booleanMutex  sync.RWMutex
	numberMutex   sync.RWMutex
	durationMutex sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		booleanTaskProgress:  make(map[uuid.UUID]map[time.Time]bool),
		numberTaskProgress:   make(map[uuid.UUID]map[time.Time]float64),
		durationTaskProgress: make(map[uuid.UUID]map[time.Time]time.Duration),
	}
}

func (r *MemoryRepository) AddTask(task Tasks.Task) error {
	r.booleanMutex.Lock()
	defer r.booleanMutex.Unlock()

	switch task.TaskType {
	case Tasks.BooleanTask:
		r.booleanTaskProgress[task.Uuid] = make(map[time.Time]bool)
	case Tasks.NumberTask:
		r.numberTaskProgress[task.Uuid] = make(map[time.Time]float64)
	case Tasks.DurationTask:
		r.durationTaskProgress[task.Uuid] = make(map[time.Time]time.Duration)
	}

	return nil
}

func (r *MemoryRepository) GetAllProgress(taskUuid uuid.UUID) (PrintableProgress, bool, error) {
	booleanProgress, found, err := r.GetBooleanProgressAll(taskUuid)
	if found {
		return booleanProgress, true, err
	}

	numberProgress, found, err := r.GetNumberProgressAll(taskUuid)
	if found {
		return numberProgress, true, err
	}

	durationProgress, found, err := r.GetDurationProgressAll(taskUuid)
	if found {
		return durationProgress, true, err
	}

	return nil, false, fmt.Errorf("progress for task %s not found", taskUuid)
}

func (r *MemoryRepository) GetBooleanProgressAll(taskUuid uuid.UUID) (*BooleanProgress, bool, error) {
	r.booleanMutex.RLock()
	defer r.booleanMutex.RUnlock()

	if progress, exists := r.booleanTaskProgress[taskUuid]; exists {
		return NewBooleanProgress(progress), true, nil
	}

	return nil, false, nil
}

func (r *MemoryRepository) GetNumberProgressAll(taskUuid uuid.UUID) (*NumberProgress, bool, error) {
	r.numberMutex.RLock()
	defer r.numberMutex.RUnlock()

	if progress, exists := r.numberTaskProgress[taskUuid]; exists {
		return NewNumberProgress(progress), true, nil
	}

	return nil, false, nil
}

func (r *MemoryRepository) GetDurationProgressAll(taskUuid uuid.UUID) (*DurationProgress, bool, error) {
	r.durationMutex.RLock()
	defer r.durationMutex.RUnlock()

	if progress, exists := r.durationTaskProgress[taskUuid]; exists {
		return NewDurationProgress(progress), true, nil
	}
	return nil, false, nil
}

func (r *MemoryRepository) GetProgressBetweenDates(taskUuid uuid.UUID, from time.Time, to time.Time) (PrintableProgress, bool, error) {
	booleanProgress, found, err := r.GetBooleanProgressBetweenDates(taskUuid, from, to)
	if found {
		return booleanProgress, true, err
	}

	numberProgress, found, err := r.GetNumberProgressBetweenDates(taskUuid, from, to)
	if found {
		return numberProgress, true, err
	}

	durationProgress, found, err := r.GetDurationProgressBetweenDates(taskUuid, from, to)
	if found {
		return durationProgress, true, err
	}

	return nil, false, fmt.Errorf("progress for task %s not found", taskUuid)
}

func (r *MemoryRepository) GetBooleanProgressBetweenDates(taskUuid uuid.UUID, from time.Time, to time.Time) (*BooleanProgress, bool, error) {
	r.booleanMutex.RLock()
	defer r.booleanMutex.RUnlock()

	if progress, exists := r.booleanTaskProgress[taskUuid]; exists {
		progressBetweenDates := make(map[time.Time]bool)
		// TODO replace with more efficient structure. b-tree? sorted slice? skip list?
		for key, value := range progress {
			if key.Compare(from) >= 0 && key.Before(to) {
				progressBetweenDates[key] = value
			}
		}
		return NewBooleanProgress(progressBetweenDates), true, nil
	}

	return nil, false, nil
}

func (r *MemoryRepository) GetNumberProgressBetweenDates(taskUuid uuid.UUID, from time.Time, to time.Time) (*NumberProgress, bool, error) {
	r.numberMutex.RLock()
	defer r.numberMutex.RUnlock()

	if progress, exists := r.numberTaskProgress[taskUuid]; exists {
		progressBetweenDates := make(map[time.Time]float64)
		// TODO replace with more efficient structure. b-tree? sorted slice? skip list?
		for key, value := range progress {
			if key.Compare(from) >= 0 && key.Before(to) {
				progressBetweenDates[key] = value
			}
		}
		return NewNumberProgress(progressBetweenDates), true, nil
	}

	return nil, false, nil
}

func (r *MemoryRepository) GetDurationProgressBetweenDates(taskUuid uuid.UUID, from time.Time, to time.Time) (*DurationProgress, bool, error) {
	r.numberMutex.RLock()
	defer r.numberMutex.RUnlock()

	if progress, exists := r.durationTaskProgress[taskUuid]; exists {
		progressBetweenDates := make(map[time.Time]time.Duration)
		// TODO replace with more efficient structure. b-tree? sorted slice? skip list?
		for key, value := range progress {
			if key.Compare(from) >= 0 && key.Before(to) {
				progressBetweenDates[key] = value
			}
		}
		return NewDurationProgress(progressBetweenDates), true, nil
	}

	return nil, false, nil
}

func (r *MemoryRepository) UpdateBooleanProgress(taskUuid uuid.UUID, date time.Time, done bool) error {
	r.booleanMutex.Lock()
	defer r.booleanMutex.Unlock()

	if progress, exists := r.booleanTaskProgress[taskUuid]; exists {
		progress[date] = done
		return nil
	}

	return fmt.Errorf("progress for boolean task %s not found", taskUuid)
}

func (r *MemoryRepository) UpdateNumberProgress(taskUuid uuid.UUID, date time.Time, value float64) error {
	r.numberMutex.Lock()
	defer r.numberMutex.Unlock()

	if progress, exists := r.numberTaskProgress[taskUuid]; exists {
		progress[date] = value
		return nil
	}

	return fmt.Errorf("progress for number task %s not found", taskUuid)
}

func (r *MemoryRepository) UpdateDurationProgress(taskUuid uuid.UUID, date time.Time, value time.Duration) error {
	r.durationMutex.Lock()
	defer r.durationMutex.Unlock()

	if progress, exists := r.durationTaskProgress[taskUuid]; exists {
		progress[date] = value
		return nil
	}

	return fmt.Errorf("progress for duration task %s not found", taskUuid)
}

func (r *MemoryRepository) RemoveTaskAndProgress(taskUuid uuid.UUID) error {
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
