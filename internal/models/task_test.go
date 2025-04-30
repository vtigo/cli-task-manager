package models

import (
	"testing"
)

func TestNewTaskList(t *testing.T) {
	list := NewTaskList("list")

	if list.Name != "list" {
		t.Errorf("Expected Task List name to be list, it is: %s", list.Name)
	}

	if len(list.Tasks) != 0 {
		t.Errorf("Expected Task List length to be 0, it is: %v", len(list.Tasks))
	}

	// TODO: Test Id
}

func TestNewTask(t *testing.T) {
	list := NewTaskList("list")

	task := NewTask(list, "task")

	if task.Name != "task" {
		t.Errorf("Expected Task name to be task, it is: %s", list.Name)
	}

	if task.Completed {
		t.Error("Expected Task to have Completed = false")
	}

	if list.Tasks[0] != *task {
		t.Errorf("Expected task to be added to the given list")
	}

	// TODO: Test Id
}

