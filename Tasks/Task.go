package Tasks

import (
	"fmt"
	"github.com/google/uuid"
)

type Task struct {
	taskType TaskType
	Uuid     uuid.UUID
	Name     string
}

func NewTask(name string) Task {
	return Task{Uuid: uuid.New(), Name: name}
}

func (t Task) String() string {
	return fmt.Sprintf("%s uuid=%s, Name=%s", t.taskType, t.Uuid, t.Name)
}

type TaskType int

func (t TaskType) String() string {
	switch t {
	case BooleanTask:
		return "BooleanTask"
	case NumberTask:
		return "NumberTask"
	case DurationTask:
		return "DurationTask"
	default:
		return "UnknownTask"
	}
}

const (
	BooleanTask TaskType = iota
	NumberTask
	DurationTask
)
