package handlers

import (
	"errors"
)

func (h *TaskHandler) HandleMarkAsCompleted(index int) error {
	if index > len(h.taskManager.Tasks) {
		return errors.New("index is out of range")
	}

	h.taskManager.Tasks[index].MarkAsCompleted()
	err := h.storage.SaveTasks(h.taskManager.Tasks)
	if err != nil {
		return err
	}

	return nil
}
