package handlers

import (
	"fmt"
)

func (h *TaskHandler) HandleListTasks() {
	fmt.Println("-- Tasks --")
	fmt.Println("-----------")
	
	cmpString := "[X]"
	notCmpString := "[ ]"
	
	if len(h.taskManager.Tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	for i, t := range h.taskManager.Tasks {
		statusString := notCmpString
		if t.Completed {
			statusString = cmpString
		}
		fmt.Printf("%v) %s %s\n", i+1, statusString, t.Name)
	}
}
