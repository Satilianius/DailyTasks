package main

type TaskRepository struct {
	tasks []Task
}

func (taskRepository *TaskRepository) AddTask(task Task) {
	taskRepository.tasks = append(taskRepository.tasks, task)
}

func (taskRepository *TaskRepository) Tasks() []Task {
	return taskRepository.tasks
}
