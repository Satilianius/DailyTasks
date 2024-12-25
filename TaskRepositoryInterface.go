package main

import "github.com/google/uuid"

type TaskRepository interface {
	Add(task Task) error
	GetByUuid(uuid uuid.UUID) (*Task, error)
	Update(updatedTask Task) error
	GetAll() ([]Task, error)
	Remove(uuid uuid.UUID) error
}
