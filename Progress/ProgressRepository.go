package Progress

import "github.com/google/uuid"

type Repository interface {
	Add(progress PrintableProgress) error
	GetByUuid(taskUuid uuid.UUID) (*PrintableProgress, error)
	Update(updatedProgress PrintableProgress) error
	GetAll() ([]PrintableProgress, error)
	Remove(taskUuid uuid.UUID) error
}
