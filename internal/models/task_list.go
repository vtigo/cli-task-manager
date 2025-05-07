package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type TaskList struct {
	Tasks []*Task
}

func NewTaskList() *TaskList {
	return &TaskList{
		Tasks: make([]*Task, 0),
	}
}

func (tl *TaskList) Json() []byte {
	jsonBytes, err := json.Marshal(tl)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return jsonBytes
}

func (tl *TaskList) AddTask(t *Task) {
	tl.Tasks = append(tl.Tasks, t)
}

func (tl *TaskList) SaveFile() {
	os.Mkdir("storage", os.ModePerm)
	path := filepath.Join("storage", fmt.Sprintf("%s.json", "tasks"))
	os.WriteFile(path, tl.Json(), 0644)
}
