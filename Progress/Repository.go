package Progress

import (
	"DailyTasks/Tasks"
	"github.com/google/uuid"
	"time"
)

type Repository interface {
	AddBooleanTask(task Tasks.Task) error
	AddNumberTask(task Tasks.Task) error
	GetPrintableProgressByUuid(taskUuid uuid.UUID) (PrintableProgress, error)
	GetBooleanByUuid(taskUuid uuid.UUID) (*BooleanProgress, bool, error)
	getNumberByUuid(taskUuid uuid.UUID) (*NumberProgress, bool, error)
	UpdateBooleanProgress(taskUuid uuid.UUID, date time.Time, done bool) error
	UpdateNumberProgress(taskUuid uuid.UUID, date time.Time, value float64) error
	Remove(taskUuid uuid.UUID) error
}
