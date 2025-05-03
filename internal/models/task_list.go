package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var TaskListId int
var TaskLists []*TaskList

type TaskList struct {
	Id    int
	Name  string
	Tasks []Task
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
	tl := &TaskList{
		Id: TaskListId,
		Name: name,
		Tasks: nil,
	}
	TaskListId++
	
	TaskLists = append(TaskLists, tl)

	return tl
}
