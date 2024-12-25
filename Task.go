package main

import (
	"fmt"
	"github.com/google/uuid"
)

type Task struct {
	uuid uuid.UUID
	name string
}

func NewTask(name string) Task {
	return Task{uuid: uuid.New(), name: name}
}

func (t Task) String() string {
	return fmt.Sprintf("uuid=%s, Name=%s", t.uuid, t.name)
}
