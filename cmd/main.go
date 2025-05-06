package main

import (
	"fmt"

	"github.com/vtigo/cli-task-manager/internal/db"
	"github.com/vtigo/cli-task-manager/internal/handlers"
	"github.com/vtigo/cli-task-manager/internal/models"
)

func main() {
	taskManager := models.NewTaskManager()
	storage := db.NewFileStorage("storage")
	handler := handlers.NewTaskHandler(taskManager, storage)

	taskLists, err := storage.LoadAllTaskLists()
	if err != nil {
		fmt.Println("Failed to load task lists: ", err)
	}

	for _, tl := range taskLists {
		taskManager.TaskLists = append(taskManager.TaskLists, tl)
	}
	handler.ListTaskLists()
}

