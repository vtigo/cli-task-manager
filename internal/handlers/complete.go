package handlers

import (
	"errors"
)

func HandleMarkAsCompleted(h *TaskHandler, index int) error {
	if index > len(h.taskManager.Tasks) {
		return errors.New("index is out of range")
	}
	//TODO: mark task as completed
	// h.taskManager.[index].MarkAsCompleted()
	// tm.HandleSaveLists()
	return nil
}
