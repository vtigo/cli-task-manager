package main

import (
	"github.com/vtigo/cli-task-manager/internal/models"
)

func main() {
	// list := models.NewTaskList("list2")
	// models.NewTask(list, "task")
	// list.SaveFile()
	models.LoadTaskLists()
	models.ListTaskLists()
}


