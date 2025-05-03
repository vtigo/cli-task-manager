package models

var taskId int
var Tasks []*Task

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool	 `json:"completed"`
}

func (t *Task) MarkAsCompleted() {
	t.Completed = true
}

func NewTask(list *TaskList, name string) *Task {
	t := &Task{
		Id: 0,
		Name: name,
		Completed: false,
	}
	taskId++
	
	Tasks = append(Tasks, t)
	list.AddTask(t)

	return t
}

