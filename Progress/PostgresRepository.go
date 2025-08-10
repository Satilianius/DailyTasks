package Progress

import (
	"DailyTasks/Tasks"
	"github.com/google/uuid"
	"time"
)

type PostgresRepository struct{}

func (PostgresRepository) AddTask(task Tasks.Task) error {
	//TODO implement me
	panic("implement me")
}

func (PostgresRepository) GetAllProgress(taskUuid uuid.UUID) (PrintableProgress, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (PostgresRepository) GetProgressBetweenDates(taskUuid uuid.UUID, from time.Time, to time.Time) (PrintableProgress, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (PostgresRepository) UpdateBooleanProgress(taskUuid uuid.UUID, date time.Time, done bool) error {
	//TODO implement me
	panic("implement me")
}

func (PostgresRepository) UpdateNumberProgress(taskUuid uuid.UUID, date time.Time, value float64) error {
	//TODO implement me
	panic("implement me")
}

func (PostgresRepository) UpdateDurationProgress(taskUuid uuid.UUID, date time.Time, value time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (PostgresRepository) RemoveTaskAndProgress(taskUuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
