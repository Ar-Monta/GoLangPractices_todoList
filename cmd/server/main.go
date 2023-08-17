package main

import (
	"fmt"
	"github.com/ArMo-Team/GoLangPractices_todoList/internal/commands"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: migrate | start")
		os.Exit(1)
	}

	commandName := os.Args[1]

	var command commands.Command

	switch commandName {
	case "migrate":
		command = &commands.MigrateCommand{}
	case "start":
		command = &commands.StartCommand{}

	default:
		fmt.Printf("Unknown command: %s\n", commandName)
		os.Exit(1)
	}

	command.Execute()
}
