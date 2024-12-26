package main

import (
	"DailyTasks/Tasks"
	"fmt"
)

func main() {
	var taskRepository Tasks.TaskRepository = Tasks.NewMemoryTaskRepository()
	task := Tasks.NewTask("testTask")

	err := taskRepository.Add(task)
	if err != nil {
		fmt.Println(err)
		return
	}

	tasks, err := taskRepository.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("tasks: %s\n", tasks)
}
