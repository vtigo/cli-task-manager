package handlers

import "fmt"

func (h *TaskHandler) HandleCreateTask(name string) error {
	_, err := h.taskManager.CreateTask(name)
	if err != nil {
		fmt.Println("Failed to create task: ", err)
		return fmt.Errorf("failed to create task list: %w", err)
	}
	
	err = h.storage.SaveTasks(h.taskManager.Tasks)
	if err != nil {
		fmt.Println("Failed to save tasks: ", err)
		return fmt.Errorf("failed to save tasks: %w", err)
	}
	
	return nil
}
