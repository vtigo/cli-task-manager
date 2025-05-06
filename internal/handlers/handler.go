package handlers

import (
	"github.com/vtigo/cli-task-manager/internal/db"
	"github.com/vtigo/cli-task-manager/internal/models"
)

type TaskHandler struct {
	taskManager *models.TaskManager
	storage     *db.FileStorage
}

func NewTaskHandler(tm *models.TaskManager, s *db.FileStorage) *TaskHandler {
	return &TaskHandler{
		taskManager: tm,
		storage:     s,
	}
}
