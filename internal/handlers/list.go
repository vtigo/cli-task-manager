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
	
	cmpString := "[X]"
	notCmpString := "[ ]"
	
	var statusString string

	for i, t := range models.Tasks { 
		if t.Completed {
			statusString = cmpString
		} else {
			statusString = notCmpString
		}
		fmt.Printf("%v) %s %s\n", i + 1, statusString, t.Name)
	}
}
