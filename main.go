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

	tasks, err := taskRepository.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	tasksWithProgress := make(map[Tasks.Task]Progress.PrintableProgress)
	for _, task := range tasks {
		printableProgress, err := progressRepository.GetPrintableProgressByUuid(task.Uuid)
		if err != nil {
			fmt.Printf("Error while getting progress for task %s:\n\n%v", task.Uuid, err)
			return
		}
		tasksWithProgress[task] = printableProgress
	}

	printProgress(tasksWithProgress)
}

func printProgress(tasksWithProgress map[Tasks.Task]Progress.PrintableProgress) {
	dates := getThisWeek()
	printHeader(dates[:])

	for task, progress := range tasksWithProgress {
		fmt.Printf("%15s:", task.Name)
		for _, date := range dates {
			fmt.Printf("%11s", progress.GetPrintableProgressAtDate(date))
		}
		fmt.Println()
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

func printHeader(dates []time.Time) {
	fmt.Printf("%15s:", "Tasks/Dates")

	for _, date := range dates {
		fmt.Printf("%11s", date.Format("02/01/2006"))
	}
	fmt.Println("")
}

func getThisWeek() [7]time.Time {
	today := today()
	var thisWeek [7]time.Time
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
