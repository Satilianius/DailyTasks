package main

import (
	"errors"
	"github.com/google/uuid"
	"sync"
)

type TaskRepository struct {
	tasks []Task
	mu    sync.Mutex // For thread safety
}

func (taskRepository *TaskRepository) AddTask(task Task) error {
	taskRepository.mu.Lock()
	defer taskRepository.mu.Unlock()

	// Check if the item with the same ID already exists
	for _, existingItem := range taskRepository.tasks {
		if existingItem.uuid == task.uuid {
			return errors.New("task with this ID already exists")
		}
	}
	taskRepository.tasks = append(taskRepository.tasks, task)
	return nil
}

func (taskRepository *TaskRepository) GetById(uuid uuid.UUID) (*Task, error) {
	taskRepository.mu.Lock()
	defer taskRepository.mu.Unlock()

	for _, task := range taskRepository.tasks {
		if task.uuid == uuid {
			return &task, nil
		}
	}

	return nil, errors.New("task not found")
}

func (taskRepository *TaskRepository) Update(updatedTask Task) error {
	taskRepository.mu.Lock()
	defer taskRepository.mu.Unlock()

	for i, task := range taskRepository.tasks {
		if task.uuid == updatedTask.uuid {
			taskRepository.tasks[i] = updatedTask
			return nil
		}
	}

	return errors.New("task not found")
}

func (taskRepository *TaskRepository) GetAll() ([]Task, error) {
	taskRepository.mu.Lock()
	defer taskRepository.mu.Unlock()

	tasksCopy := make([]Task, len(taskRepository.tasks))
	copy(tasksCopy, taskRepository.tasks)
	return tasksCopy, nil
}

func (taskRepository *TaskRepository) RemoveTask(uuid uuid.UUID) error {
	taskRepository.mu.Lock()
	defer taskRepository.mu.Unlock()

	for i, task := range taskRepository.tasks {
		if task.uuid == uuid {
			taskRepository.tasks = append(taskRepository.tasks[:i], taskRepository.tasks[i+1:]...)
			return nil
		}
	}

	return nil
}
