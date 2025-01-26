package Progress

import (
	"DailyTasks/Tasks"
	"github.com/google/uuid"
	"time"
)

type Repository interface {
	AddTask(task Tasks.Task) error
	GetPrintableProgressByUuid(taskUuid uuid.UUID) (PrintableProgress, error)

	GetBooleanByUuid(taskUuid uuid.UUID) (*BooleanProgress, bool, error)
	GetNumberByUuid(taskUuid uuid.UUID) (*NumberProgress, bool, error)
	GetDurationByUuid(taskUuid uuid.UUID) (*DurationProgress, bool, error)

	UpdateBooleanProgress(taskUuid uuid.UUID, date time.Time, done bool) error
	UpdateNumberProgress(taskUuid uuid.UUID, date time.Time, value float64) error
	UpdateDurationProgress(taskUuid uuid.UUID, date time.Time, value time.Duration) error

	Remove(taskUuid uuid.UUID) error
}
