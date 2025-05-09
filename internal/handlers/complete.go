package handlers

import (
	"fmt"
)

func (h *TaskHandler) HandleMarkAsCompleted(index int) error {
	if index < 0 || index >= len(h.taskManager.Tasks) {
		return fmt.Errorf("index is out of range")
	}

	h.taskManager.Tasks[index].MarkAsCompleted()

	err := h.storage.SaveTasks(h.taskManager.Tasks)
	if err != nil {
		return fmt.Errorf("failed to save tasks: %w", err)
	}

	return nil
}
