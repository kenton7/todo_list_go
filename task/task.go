package task

import "time"

type Task struct {
	Header string
	Description string
	Status Status
	CreatedAt time.Time
	FinishedAt time.Time
}

type Status string 

const (
	Undone Status = "Не выполнена"
	Done Status = "Выполнена"
)