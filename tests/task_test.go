package tests

import (
	"testing"

	"github.com/vtigo/cli-task-manager/internal/models"
)


func TestNewTaskManager(t *testing.T) {
	taskManager := models.NewTaskManager()
	
	if len(taskManager.Tasks) != 0 {
		t.Errorf("Expected 0 tasks upon initialization, got %d", len(taskManager.Tasks))
	}
}

func TestNewTask(t *testing.T) {
	taskManager := models.NewTaskManager()

	task, err := taskManager.CreateTask("task")
	if err != nil {
		t.Errorf("Expected no error upon task creation, got %d", err)
	}

	if task.Name != "task" {
		t.Errorf("Expected task name to be task, got %s", task.Name)
	}

	if task.Completed {
		t.Error("Expected task to have completed false")
	}
}

func TestToggleTaskCompleteness(t *testing.T) {
	taskManager := models.NewTaskManager()

	task, err := taskManager.CreateTask("task")
	if err != nil {
		t.Errorf("Expected no error upon task creation, got %d", err)
	}

	task.ToggleCompleteness()
	if !task.Completed {
		t.Error("Expected task to be completed after toggling one time")
	}

	task.ToggleCompleteness()
	if task.Completed {
		t.Error("Expected task to not be completed after toggling the second time")
	}
}
