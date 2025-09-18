package userEvents

import "time"

type Event struct {
	Action EventType
	Text string
	ErrorDescription string 
	CreatedAt time.Time
}

type EventType string

const (
	AddEvent EventType = "Добавление задачи"
	DeleteEvent EventType = "Удаление задачи"
	MarkAsDoneEvent EventType = "Отметка задачи как выполненной"
	UserInputEvent EventType = "Ввод пользователя"
)