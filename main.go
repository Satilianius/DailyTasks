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
		printableProgress, found, err := progressRepository.GetAllProgress(task.Uuid)
		if err != nil {
			fmt.Printf("Error while getting progress for task %s:\n\n%v", task.Uuid, err)
			continue
		}
		if !found {
			fmt.Printf("Progress not found for task %s\n", task.Uuid)
			continue
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
	booleanTask := Tasks.NewTask(Tasks.BooleanTask, "booleanTask")
	numberTask := Tasks.NewTask(Tasks.NumberTask, "numberTask")
	durationTask := Tasks.NewTask(Tasks.DurationTask, "durationTask")

	err := taskRepository.Add(booleanTask)
	if err != nil {
		return err
	}
	err = taskRepository.Add(numberTask)
	if err != nil {
		return err
	}
	err = taskRepository.Add(durationTask)
	if err != nil {
		return err
	}

	err = progressRepository.AddTask(booleanTask)
	if err != nil {
		return err
	}
	err = progressRepository.AddTask(numberTask)
	if err != nil {
		return err
	}
	err = progressRepository.AddTask(durationTask)
	if err != nil {
		return err
	}

	err = progressRepository.UpdateBooleanProgress(booleanTask.Uuid, today(), true)
	if err != nil {
		return err
	}
	err = progressRepository.UpdateNumberProgress(numberTask.Uuid, today(), 42)
	if err != nil {
		return err
	}
	err = progressRepository.UpdateDurationProgress(durationTask.Uuid, today(), time.Minute)
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
	return getStartOfDay(time.Now())
}

func getStartOfDay(day time.Time) time.Time {
	return day.Truncate(24 * time.Hour)
}
