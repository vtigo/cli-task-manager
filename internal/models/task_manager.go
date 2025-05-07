package models

type TaskManager struct {
	nextTaskId     int
	Tasks          []*Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		nextTaskId:     0,
		Tasks:          []*Task{},
	}
}

func (tm *TaskManager) CreateTask(name string) (*Task, error) {
	t := &Task{
		Id: tm.nextTaskId,
		Name: name,
		Completed: false,
	}

	tm.nextTaskId++
	tm.Tasks = append(tm.Tasks, t)

	return t, nil
}
