package main

import (
	"fmt"
	//"strconv"

	"main.go/commandsList"
	// "main.go/storage"
	// "main.go/userEvents"
)

func main() {
	fmt.Println("Добро пожаловать в ваш список дел!")
	//commands.CommandsDescription()

	// fmt.Print("Введите команду: ")
	// var userCommand int 
	// fmt.Scan(&userCommand)
	// storage.RegisterEvent(userEvents.UserInputEvent, strconv.Itoa(userCommand), "")
	commands.CommandsRoute()
}
