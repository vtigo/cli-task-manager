package models

import (
	"fmt"
	"os"
	"encoding/json"
)

var newTaskId = 0

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool	 `json:"completed"`
}

type TaskList struct {
	Name  string
	Tasks []Task
}

func (t *Task) MarkAsCompleted() {
	t.Completed = true
}

func (tl *TaskList) Json() []byte {
	jsonBytes, err := json.Marshal(tl)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return jsonBytes
}

func (tl *TaskList) AddTask(t *Task) {
	tl.Tasks = append(tl.Tasks, *t)
}

func (tl *TaskList) SaveFile() {
	os.WriteFile(fmt.Sprintf("./%s.json", tl.Name), tl.Json(), os.ModePerm)
}

func NewTaskList(name string) *TaskList {
	return &TaskList{
		Name: name,
		Tasks: nil,
	}
}

func NewTask(list *TaskList, name string) *Task {
	task := &Task{
		Id: newTaskId,
		Name: name,
		Completed: false,
	}
	list.AddTask(task)
	newTaskId++
	return task
}
