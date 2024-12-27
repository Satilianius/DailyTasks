package Tasks

import (
	"fmt"
	"github.com/google/uuid"
)

type Task struct {
	Uuid uuid.UUID
	Name string
}

func NewTask(name string) Task {
	return Task{Uuid: uuid.New(), Name: name}
}

func (t Task) String() string {
	return fmt.Sprintf("uuid=%s, Name=%s", t.Uuid, t.Name)
}
