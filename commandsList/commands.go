package commands

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/k0kubun/pp"
	"main.go/storage"
	"main.go/task"
	"main.go/userEvents"
)

func CommandsRoute() {
	var userCommand int
	reader := bufio.NewReader(os.Stdin)
	for {
		CommandsDescription()
		fmt.Print("Введите команду: ")
		fmt.Scan(&userCommand)
		storage.RegisterEvent(userEvents.UserInputEvent, strconv.Itoa(userCommand), "")
		switch userCommand {
		case 1:
			fmt.Print("Введите заголовок задачи (одно слово): ")
			header, _ := reader.ReadString('\n')
			header = strings.TrimSpace(header)
			headerWords := strings.Fields(header)
			if len(headerWords) > 1 {
				fmt.Println("ОШИБКА: Заголовок задачи должен быть из одного слова")
				storage.RegisterEvent(userEvents.UserInputEvent, header, "Заголовок задачи должен быть из одного слова")
				break
			}
			storage.RegisterEvent(userEvents.UserInputEvent, header, "")
			fmt.Print("Введите описание задачи (одно или несколько слов): ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)
			storage.RegisterEvent(userEvents.UserInputEvent, description, "")
			add(header, description)
		case 2:
			tasks := storage.GetAllTasks()
			if len(tasks) == 0 {
				pp.Println("Список задач пуст")
				break
			}
			pp.Println("Список задач:")
			pp.Println(tasks)
		case 3:
			fmt.Print("Введите заголовок задачи, которую хотите удалить: ")
			taskHeader, _ := reader.ReadString('\n')
			taskHeader = strings.TrimSpace(taskHeader)
			storage.RegisterEvent(userEvents.UserInputEvent, taskHeader, "")
			storage.DeleteFromStorage(taskHeader)
		case 4:
			fmt.Print("Введите заголовок задачи, которую хотите отметить как выполненную: ")
			taskHeader, _ := reader.ReadString('\n')
			taskHeader = strings.TrimSpace(taskHeader)
			storage.RegisterEvent(userEvents.UserInputEvent, taskHeader, "")
			storage.MarkAsDone(taskHeader)
		case 5:
			events := storage.GetAllEvents()
			pp.Println(events)
		case 6:
			help()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Команда не распознана, попробуйте еще раз")
		}
	}
}

func CommandsDescription() {
	fmt.Println("1 - добавить задачу")
	fmt.Println("2 - вывести список задач")
	fmt.Println("3 - удалить задачу")
	fmt.Println("4 - отметить задачу как выполненную")
	fmt.Println("5 - вывести список событий")
	fmt.Println("6 - вывести список команд")
	fmt.Println("0 - завершить работу")
}

func add(header string, description string) {
	task := task.Task{
		Header: 	header,
		Description: description,
		Status: 	task.Undone,
		CreatedAt: 	time.Now(),
	}
	storage.PutToStorage(task)
	storage.RegisterEvent(userEvents.AddEvent, header, "")
}

func help() {
	fmt.Println("Выберите команду, о которой хотите узнать подробнее:")
	CommandsDescription()
	fmt.Print("Введите номер команды: ")
	var command int
	fmt.Scan(&command)
	storage.RegisterEvent(userEvents.UserInputEvent, strconv.Itoa(command), "")

	switch command {
	case 1:
		addDescription := `add - добавить задачу в список дел. 
После ввода команды Вам будет предложено ввести заголовок и описание задачи.
Заголовок задачи должен быть из одного слова, описание задачи может быть из нескольких слов.`
		pp.Println(addDescription)
	case 2:
		listDescription := `list - вывести список всех задач.
После ввода команды Вам будет выведен список всех задач с их статусами (выполнена/не выполнена).`
		pp.Println(listDescription)
	case 3:
		delDescription := `del - удалить задачу из списка дел.
После ввода команды Вам нужно будет ввести заголовок задачи, которую Вы хотите удалить.`
		pp.Println(delDescription)
	case 4:
		doneDescription := `done - отметить задачу как выполненную.
После ввода команды Вам нужно будет ввести заголовок задачи, которую Вы хотите отметить как выполненную.`
		pp.Println(doneDescription)
	case 5:
		pp.Println("events - вывести список всех событий")
	case 0:
		pp.Println("exit - завершить работу программы")
	default:
		pp.Println("Команда не распознана, попробуйте еще раз")
		help()
	}
}
