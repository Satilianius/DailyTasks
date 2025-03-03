package Tasks

import (
	"errors"
	"github.com/google/uuid"
	"slices"
	"sync"
)

type MemoryTaskRepository struct {
	tasks []Task
	mu    sync.RWMutex
}

func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{
		tasks: []Task{},
	}
}

func (r *MemoryTaskRepository) Add(task Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if the item with the same ID already exists
	idx := slices.IndexFunc(r.tasks, func(t Task) bool {
		return t.Uuid == task.Uuid
	})

	if idx != -1 {
		return errors.New("task with this ID already exists")
	}
	r.tasks = append(r.tasks, task)
	return nil
}

func (r *MemoryTaskRepository) GetByUuid(uuid uuid.UUID) (*Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	idx := slices.IndexFunc(r.tasks, func(t Task) bool {
		return t.Uuid == uuid
	})

	if idx == -1 {
		return nil, errors.New("task not found")
	}

	task := r.tasks[idx]
	return &task, nil
}

func (r *MemoryTaskRepository) Update(updatedTask Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	idx := slices.IndexFunc(r.tasks, func(t Task) bool {
		return t.Uuid == updatedTask.Uuid
	})

	if idx == -1 {
		return errors.New("task not found")
	}

	r.tasks[idx] = updatedTask
	return nil
}

func (r *MemoryTaskRepository) GetAll() ([]Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return slices.Clone(r.tasks), nil
}

func (r *MemoryTaskRepository) Remove(uuid uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	idx := slices.IndexFunc(r.tasks, func(t Task) bool {
		return t.Uuid == uuid
	})

	if idx == -1 {
		return nil
	}

	r.tasks = slices.Delete(r.tasks, idx, idx+1)
	return nil
}
