package storage

import (
	"errors"
	"fmt"
	"time"

	"main.go/task"
	"main.go/userEvents"
)

var storage = make(map[string]task.Task)
var events = make(map[string]userEvents.Event)

func PutToStorage(t task.Task) {
	storage[t.Header] = t
	fmt.Println("Задача успешно добавлена!")
}

func GetAllTasks() map[string]task.Task {
	return storage
}

func DeleteFromStorage(taskHeader string) error {
	_, ok := storage[taskHeader]
	if !ok {
		fmt.Println("ОШИБКА: Задача с таким заголовком не найдена")
		RegisterEvent(userEvents.DeleteEvent, taskHeader, "ОШИБКА: Задача с таким заголовком не найдена")
		return errors.New("ОШИБКА: Задача с таким заголовком не найдена")
	}
	delete(storage, taskHeader)
	fmt.Println("Задача успешно удалена!")
	RegisterEvent(userEvents.DeleteEvent, taskHeader, "")
	return nil
}

func MarkAsDone(taskHeader string) error {
	t, ok := storage[taskHeader]
	if !ok {
		fmt.Println("ОШИБКА: Задача с таким заголовком не найдена")
		RegisterEvent(userEvents.MarkAsDoneEvent, taskHeader, "Задача с таким заголовком не найдена")
		return errors.New("ОШИБКА: Задача с таким заголовком не найдена")
	}
	t.Status = task.Done
	t.FinishedAt = time.Now()
	storage[taskHeader] = t
	fmt.Println("Задача успешно отмечена как выполненная!")
	RegisterEvent(userEvents.MarkAsDoneEvent, taskHeader, "")
	return nil
}

func RegisterEvent(action userEvents.EventType, text string, description string) {
	event := userEvents.Event{
		Action: action,
		Text: text,
		ErrorDescription: description,
		CreatedAt: time.Now(),
	}
	events[text] = event
}

func GetAllEvents() map[string]userEvents.Event {
	return events
}
