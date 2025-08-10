package Tasks

import "github.com/google/uuid"

type PostgresRepository struct {
}

func (PostgresRepository) Add(task Task) error {
	//TODO implement me
	panic("implement me")
}

func (PostgresRepository) GetByUuid(uuid uuid.UUID) (*Task, error) {
	//TODO implement me
	panic("implement me")
}

func (PostgresRepository) Update(updatedTask Task) error {
	//TODO implement me
	panic("implement me")
}

func (PostgresRepository) GetAll() ([]Task, error) {
	//TODO implement me
	panic("implement me")
}

func (PostgresRepository) Remove(uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
