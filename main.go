package main

import (
	"DailyTasks/Progress"
	"DailyTasks/Tasks"
	"fmt"
	"time"
)

func main() {
	taskRepository := Tasks.TaskRepository(Tasks.NewMemoryTaskRepository())
	progressRepository := Progress.Repository(Progress.NewMemoryRepository())

	err := fillRepositories(taskRepository, progressRepository)
	if err != nil {
		fmt.Println(err)
		return
	}

	printHeader()

	tasks, err := taskRepository.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, task := range tasks {
		printableProgress, err := progressRepository.GetByUuid(task.Uuid)
		if err != nil {
			fmt.Printf("Error while getting progress for task %s:\n\n%v", task.Uuid, err)
			return
		}
		fmt.Printf("%15s: %s\n", task.Name, printableProgress.GetPrintableProgressAtDate(today()))
	}
}

func fillRepositories(taskRepository Tasks.TaskRepository, progressRepository Progress.Repository) error {
	booleanTask := Tasks.NewTask("booleanTask")
	numberTask := Tasks.NewTask("numberTask")

	err := taskRepository.Add(booleanTask)
	if err != nil {
		return err
	}
	err = taskRepository.Add(numberTask)
	if err != nil {
		return err
	}

	err = progressRepository.AddBooleanTask(booleanTask)
	if err != nil {
		return err
	}
	err = progressRepository.AddNumberTask(numberTask)
	if err != nil {
		return err
	}
	return nil
}

func printHeader() {
	fmt.Printf("%15s:", "Tasks/Dates")
	var dates = getThisWeek()
	for _, date := range dates {
		fmt.Printf("%11s", date.Format("02/01/2006"))
	}
	fmt.Println("")
}

func getThisWeek() [7]time.Time {
	today := today().AddDate(0, 0, 2)
	thisWeek := [7]time.Time{}
	for weekdayIndex := 0; weekdayIndex < 7; weekdayIndex++ {
		// Adjust Weekday to make Monday 0 (0 is Sunday by default)
		adjustedWeekday := (int(today.Weekday()) + 6) % 7
		daysToAdd := weekdayIndex - adjustedWeekday
		thisWeek[weekdayIndex] = today.AddDate(0, 0, daysToAdd)
	}

	return thisWeek
}

func today() time.Time {
	return time.Now().Truncate(24 * time.Hour)
}
