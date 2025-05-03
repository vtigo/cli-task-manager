package tests

import (
	"testing"

	"github.com/vtigo/cli-task-manager/internal/models"
)

func TestNewTaskList(t *testing.T) {
	list := models.NewTaskList("list")

	if list.Name != "list" {
		t.Errorf("Expected Task List name to be list, it is: %s", list.Name)
	}

	if len(list.Tasks) != 0 {
		t.Errorf("Expected Task List length to be 0, it is: %v", len(list.Tasks))
	}

	if list.Id != 0 {
		t.Errorf("Expected Task List id to be 0, it is: %v", list.Id)
	}
}

func TestNewTask(t *testing.T) {
	list := models.NewTaskList("list")
	task := models.NewTask(list, "task")

	if task.Name != "task" {
		t.Errorf("Expected Task name to be task, it is: %s", task.Name)
	}

	if task.Completed {
		t.Error("Expected Task to have Completed = false")
	}

	if list.Tasks[0] != *task {
		t.Errorf("Expected Task to be added to the given list")
	}

	if task.Id != 0 {
		t.Errorf("Expected Task id to be 0, it is: %v", task.Id)
	}
}

