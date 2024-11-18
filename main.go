package main

import (
	"fmt"
)

func printUsage() {
	fmt.Println("")
	fmt.Println(colorizeOutput("Usage: todo <command> [arguments]"))
	fmt.Println("\nCommands:")
	fmt.Println(printCommand("add", "Add a todo", 2))
	fmt.Println(printCommand("edit", "Edit a todo", 2))
	fmt.Println(printCommand("delete", "Delete a todo", 1))
	fmt.Println(printCommand("list", "List all todos", 2))
	fmt.Println(printCommand("complete", "Mark a todo as complete", 1))
	fmt.Println(printCommand("incomplete", "Mark a todo as incomplete", 1))
	fmt.Println(printCommand("help", "Print usage/help", 2))
}

func main() {
	printUsage()
}
