package handlers

import (
	"fmt"
)

func (h *TaskHandler) HandleToggleTaskCompleteness(index int) error {
	// check if the index is valid
	if index < 0 || index >= len(h.taskManager.Tasks) {
		return fmt.Errorf("index is out of range")
	}
	
	// change the task state to !completed
	h.taskManager.Tasks[index].ToggleCompleteness()
	
	// save the new state to storage
	err := h.storage.SaveTasks(h.taskManager.Tasks)
	if err != nil {
		return fmt.Errorf("failed to save tasks: %w", err)
	}

	return nil
}
