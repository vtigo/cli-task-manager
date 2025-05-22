package handlers

import (
	"fmt"
	"slices"
)

func (h *TaskHandler) HandleDeleteTask(index int) error {
	// check if the index is valid
	if index < 0 || index >= len(h.taskManager.Tasks) {
		return fmt.Errorf("index out of range")
	}
	
	// remove the task from the state
	h.taskManager.Tasks = slices.Delete(h.taskManager.Tasks, index, index + 1)
	
	// save the new state to storage
	err := h.storage.SaveTasks(h.taskManager.Tasks)
	if err != nil {
		return fmt.Errorf("failed to save tasks: %w", err)
	}

	return nil
}

