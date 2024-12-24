package main

import (
	"fmt"
)

func main() {
	taskRepository := new(TaskRepository)
	var task = Task{name: "testTask"}

	taskRepository.AddTask(task)
	tasks := taskRepository.Tasks()
	fmt.Printf("tasks: %s\n", tasks)
}
