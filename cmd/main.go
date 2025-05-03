package main

import (
	"flag"
	"fmt"

	"github.com/vtigo/cli-task-manager/internal/db"
	"github.com/vtigo/cli-task-manager/internal/handlers"
)

func main() {
	db.HandleLoadTaskLists()
	
	shouldListTasks := flag.Bool("la", false, "Lists all tasks")
	markTask := flag.Int("c", -1, " Mark task as completed")
	flag.Parse()
	
	if *shouldListTasks {
		handlers.HandleListTasks()
	}
	
	if *markTask >= 0 {
		err := handlers.HandleMarkAsCompleted(*markTask - 1)
		if err != nil {
			fmt.Println("Failed to mark task as completed: ", err)
		}
	}
}
