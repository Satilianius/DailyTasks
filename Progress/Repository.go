package Progress

import (
	"DailyTasks/Tasks"
	"github.com/google/uuid"
	"time"
)

type Repository interface {
	AddTask(task Tasks.Task) error

	GetAllProgress(taskUuid uuid.UUID) (PrintableProgress, bool, error)
	GetProgressBetweenDates(taskUuid uuid.UUID, from time.Time, to time.Time) (PrintableProgress, bool, error)

	UpdateBooleanProgress(taskUuid uuid.UUID, date time.Time, done bool) error
	UpdateNumberProgress(taskUuid uuid.UUID, date time.Time, value float64) error
	UpdateDurationProgress(taskUuid uuid.UUID, date time.Time, value time.Duration) error

	Remove(taskUuid uuid.UUID) error
}
