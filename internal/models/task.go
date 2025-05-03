package models

import "fmt"

var TaskId int
var Tasks []*Task

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool	 `json:"completed"`
}

func (t *Task) MarkAsCompleted() {
	t.Completed = true
}

func NewTask(list *TaskList, name string) *Task {
	if list == nil {
		fmt.Println("Task List do not exist.")
		return &Task{}
	}

	t := &Task{
		Id: TaskId,
		Name: name,
		Completed: false,
	}
	TaskId++
	
	Tasks = append(Tasks, t)
	list.AddTask(t)

	return t
}

