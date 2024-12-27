package main

import (
	"DailyTasks/Progress"
	"DailyTasks/Tasks"
	"fmt"
	"time"
)

func main() {
	taskRepository := Tasks.TaskRepository(Tasks.NewMemoryTaskRepository())
	testTask1 := Tasks.NewTask("testTask1")
	testTask2 := Tasks.NewTask("testTask2")

	err := taskRepository.Add(testTask1)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = taskRepository.Add(testTask2)
	if err != nil {
		fmt.Println(err)
		return
	}

	progressRepository := Progress.Repository(Progress.NewMemoryRepository())
	err = progressRepository.AddBooleanTask(testTask1)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = progressRepository.AddNumberTask(testTask2)
	if err != nil {
		fmt.Println(err)
		return
	}

	tasks, err := taskRepository.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	printHeader()

	for _, task := range tasks {
		printableProgress, err := progressRepository.GetByUuid(task.Uuid)
		if err != nil {
			fmt.Printf("Error while getting progress for task %s:\n\n%v", task.Uuid, err)
			return
		}
		fmt.Printf("%15s: %s\n", task.Name, printableProgress.GetPrintableProgressAtDate(today()))
	}
}

func printHeader() {
	fmt.Printf("%15s:", "Tasks/Dates")
	var dates = getThisWeek()
	for _, date := range dates {
		fmt.Printf("%11s", date.Format("02/01/2006"))
	}
	fmt.Println("")
}

func getThisWeek() []time.Time {
	today := today()
	// TODO add other week days
	return []time.Time{today}
}

func today() time.Time {
	return time.Now().Truncate(24 * time.Hour)
}
