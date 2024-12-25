package main

import (
	"fmt"
)

func main() {
	taskRepository := new(TaskRepository)
	var task = NewTask("testTask")

	err := taskRepository.AddTask(task)
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
