package handlers

import "fmt"

func (h *TaskHandler) HandleCreateTask(name string) bool {
	_, err := h.taskManager.CreateTask(name)
	if err != nil {
		fmt.Println("Failed to create task: ", err)
		return false
	}
	
	err = h.storage.SaveTasks(h.taskManager.Tasks)
	if err != nil {
		fmt.Println("Failed to save tasks: ", err)
		return false
	}
	
	return true
}
