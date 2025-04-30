package main

import(
	"github.com/vtigo/cli-task-manager/internal/models"
)

func main() {
	list := models.NewTaskList("list")
	models.NewTask(list, "task")
	list.SaveFile()
}


