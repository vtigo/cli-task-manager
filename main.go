package main

import (
	"fmt"
	"github.com/vtigo/cli-task-manager/internal/models"
)

func main() {
	// list1 := models.NewTaskList("list_01")
	// list2 := models.NewTaskList("list_02")
	//
	// models.NewTask(list1, "task_01")
	// models.NewTask(list1, "task_02")
	//
	// models.NewTask(list2, "task_3")
	// models.NewTask(list2, "task_4")
	//
	// list1.SaveFile()
	// list2.SaveFile()

	models.LoadTaskLists()
	models.ListTasks()
	task := models.GetTaskByName("task_01")
	fmt.Println(task.Name)
}


