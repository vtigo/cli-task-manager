package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var newTaskId = 0

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool	 `json:"completed"`
}

func (t *Task) MarkAsCompleted() {
	t.Completed = true
}

type TaskList struct {
	Name  string
	Tasks []Task
}

func NewTaskList(name string) *TaskList {
	return &TaskList{
		Name: name,
		Tasks: nil,
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
	tl.Tasks = append(tl.Tasks, *t)
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

func main() {
	taskList := NewTaskList("list1")
	NewTask(taskList, "task1")
	NewTask(taskList, "task2")
	taskList2 := NewTaskList("list2")
	NewTask(taskList2, "task3")
	os.WriteFile(fmt.Sprintf("./%s.json", taskList.Name), taskList.Json(), os.ModePerm)
	os.WriteFile(fmt.Sprintf("./%s.json", taskList2.Name), taskList.Json(), os.ModePerm)
}

