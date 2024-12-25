package main

import (
	"fmt"
)

func main() {
	var taskRepository TaskRepository = NewMemoryTaskRepository()
	task := NewTask("testTask")

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
