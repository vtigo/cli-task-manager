package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool	 `json:"completed"`
}

type TaskList struct {
	Id    int
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
	os.Mkdir("storage", os.ModePerm)
	path := filepath.Join("storage", fmt.Sprintf("%s.json", tl.Name))
	os.WriteFile(path, tl.Json(), 0644)
}

func NewTaskList(name string) *TaskList {
	// TODO: increment id
	return &TaskList{
		Id: 0,
		Name: name,
		Tasks: nil,
	}
}

func NewTask(list *TaskList, name string) *Task {
	// TODO: increment id
	task := &Task{
		Id: 0,
		Name: name,
		Completed: false,
	}
	list.AddTask(task)
	return task
}
