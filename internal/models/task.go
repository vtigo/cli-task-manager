package models

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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

var TaskLists []*TaskList

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

func LoadTaskLists() []*TaskList {
	TaskLists = []*TaskList{}

	root := os.DirFS("./storage")
	taskListFiles, err := fs.Glob(root, "*.json")
	if err != nil {
		fmt.Println("Failed to read storage directory")
		return TaskLists
	}

	for i := range taskListFiles {
		path := filepath.Join("storage", taskListFiles[i])
		file, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("failed to read file:", err)
			continue
		}
	
		name := strings.Split(taskListFiles[i], ".")[0]
		tl := NewTaskList(name)

		err = json.Unmarshal(file, &tl)
		if err != nil {
			fmt.Println("failed to unmarshal task list:", err)
			continue
		}

		TaskLists = append(TaskLists, tl)
	}

	return TaskLists
}

func ListTaskLists() {
	fmt.Println("-- Task Lists --")
	fmt.Println("----------------")

	if len(TaskLists) == 0 {
		fmt.Println("No Task Lists found")
		return
	}

	for i, tl := range TaskLists {
		fmt.Printf("%v - %s\n", i + 1, tl.Name)
	}
}
