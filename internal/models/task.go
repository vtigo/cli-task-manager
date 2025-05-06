package models

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool	 `json:"completed"`
}

func (t *Task) MarkAsCompleted() {
	t.Completed = true
}
