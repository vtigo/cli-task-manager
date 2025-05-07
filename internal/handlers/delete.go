package handlers

import (
	"fmt"
	"slices"
)

func (h *TaskHandler) HandleDeleteTask(index int) error {
	if index < 0 || index >= len(h.taskManager.Tasks) {
		return fmt.Errorf("index out of range")
	}

	h.taskManager.Tasks = slices.Delete(h.taskManager.Tasks, index, index + 1)
	
	err := h.storage.SaveTasks(h.taskManager.Tasks)
	if err != nil {
		return fmt.Errorf("failed to save tasks: %w", err)
	}

	return nil
}

