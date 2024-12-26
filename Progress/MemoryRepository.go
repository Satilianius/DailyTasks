package Progress

import (
	"github.com/google/uuid"
	"sync"
)

type MemoryRepository struct {
	taskProgress map[uuid.UUID]PrintableProgress
	mu           sync.Mutex // For thread safety
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		taskProgress: make(map[uuid.UUID]PrintableProgress),
	}
}

func (r *MemoryRepository) Add(progress PrintableProgress) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	//TODO implement me
	panic("implement me")
}

func (r *MemoryRepository) GetByUuid(taskUuid uuid.UUID) (*PrintableProgress, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	//TODO implement me
	panic("implement me")
}

func (r *MemoryRepository) Update(updatedProgress PrintableProgress) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	//TODO implement me
	panic("implement me")
}

func (r *MemoryRepository) GetAll() ([]PrintableProgress, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	//TODO implement me
	panic("implement me")
}

func (r *MemoryRepository) Remove(taskUuid uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	//TODO implement me
	panic("implement me")
}
