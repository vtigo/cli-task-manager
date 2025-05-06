package handlers

import (
	"fmt"
)

func (h *TaskHandler) ListTaskLists() {
	fmt.Println("-- Task Lists --")
	fmt.Println("----------------")

	if len(h.taskManager.TaskLists) == 0 {
		fmt.Println("No Task Lists found")
		return
	}

	for i, tl := range h.taskManager.TaskLists {
		fmt.Printf("%v - %s\n", i+1, tl.Name)
	}
}

func (h *TaskHandler) ListTasks() {
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
