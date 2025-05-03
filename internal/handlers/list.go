package handlers

import (
	"fmt"
	"github.com/vtigo/cli-task-manager/internal/models"
)

func HandleListTaskLists() {
	fmt.Println("-- Task Lists --")
	fmt.Println("----------------")

	if len(models.TaskLists) == 0 {
		fmt.Println("No Task Lists found")
		return
	}

	for i, tl := range models.TaskLists {
		fmt.Printf("%v - %s\n", i + 1, tl.Name)
	}
}

func HandleListTasks() {
	fmt.Println("-- Tasks --")
	fmt.Println("-----------")

	for i, t := range models.Tasks {
		fmt.Printf("%v - %s\n", i + 1, t.Name)
	}
}
