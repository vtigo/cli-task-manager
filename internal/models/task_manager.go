package models

type TaskManager struct {
	nextTaskId     int
	nextTaskListId int
	Tasks          []*Task
	TaskLists      []*TaskList
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		nextTaskId:     0,
		nextTaskListId: 0,
		Tasks:          []*Task{},
		TaskLists:      []*TaskList{},
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

func (tm *TaskManager) CreateTaskList(name string) (*TaskList, error) {
	tl := &TaskList{
		Id: tm.nextTaskListId,
		Name: name,
		Tasks: []Task{},
	}

	tm.nextTaskListId++
	tm.TaskLists = append(tm.TaskLists, tl)

	return tl, nil
}
