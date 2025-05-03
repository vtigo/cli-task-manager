package handlers

import (
	"errors"

	"github.com/vtigo/cli-task-manager/internal/db"
	"github.com/vtigo/cli-task-manager/internal/models"
)

func HandleMarkAsCompleted(index int) error {
	if index > len(models.Tasks) {
		return errors.New("index is out of range")
	}
	models.Tasks[index].MarkAsCompleted()
	db.HandleSaveLists()
	return nil
}
