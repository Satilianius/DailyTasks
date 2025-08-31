package Tasks

import (
	"fmt"

	"github.com/google/uuid"
)

type Task struct {
	TaskType TaskType
	Uuid     uuid.UUID
	Name     string
}

func NewTask(taskType TaskType, name string) Task {
	return Task{TaskType: taskType, Uuid: uuid.New(), Name: name}
}

func (t Task) String() string {
	return fmt.Sprintf("%s uuid=%s, Name=%s", t.TaskType, t.Uuid, t.Name)
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
