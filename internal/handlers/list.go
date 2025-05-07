package handlers

import (
	"fmt"
)

const (
	completedTaskString = "[X]"
	uncompletedTaskString = "[ ]"
)

func (h *TaskHandler) HandleListTasks() {
	fmt.Println("-- Tasks --")
	fmt.Println("-----------")
	
	
	if len(h.taskManager.Tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	for i, t := range h.taskManager.Tasks {
		statusString := uncompletedTaskString
		if t.Completed {
			statusString = completedTaskString
		}
		fmt.Printf("%v) %s %s\n", i+1, statusString, t.Name)
	}
}
