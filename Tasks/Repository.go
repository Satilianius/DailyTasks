package Tasks

import "github.com/google/uuid"

type Repository interface {
	Add(task Task) error
	GetByUuid(uuid uuid.UUID) (*Task, error)
	Update(updatedTask Task) error
	GetAll() ([]Task, error)
	Remove(uuid uuid.UUID) error
}
