package Tasks

import (
	"errors"
	"github.com/google/uuid"
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
	for _, existingItem := range r.tasks {
		if existingItem.Uuid == task.Uuid {
			return errors.New("task with this ID already exists")
		}
	}
	r.tasks = append(r.tasks, task)
	return nil
}

func (r *MemoryTaskRepository) GetByUuid(uuid uuid.UUID) (*Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, task := range r.tasks {
		if task.Uuid == uuid {
			return &task, nil
		}
	}

	return nil, errors.New("task not found")
}

func (r *MemoryTaskRepository) Update(updatedTask Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, task := range r.tasks {
		if task.Uuid == updatedTask.Uuid {
			r.tasks[i] = updatedTask
			return nil
		}
	}

	return errors.New("task not found")
}

func (r *MemoryTaskRepository) GetAll() ([]Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasksCopy := make([]Task, len(r.tasks))
	copy(tasksCopy, r.tasks)
	return tasksCopy, nil
}

func (r *MemoryTaskRepository) Remove(uuid uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, task := range r.tasks {
		if task.Uuid == uuid {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}

	return nil
}
