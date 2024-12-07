package main

import (
	"fmt"
	"os"
	"strings"
)

func printUsage() {
	fmt.Println("")
	fmt.Println(formatWithCyan("Usage: todo <command> [arguments]"))
	fmt.Println("\nCommands:")
	fmt.Println(printCommand("add\t\t<name>", "Add a todo", 2))
	fmt.Println(printCommand("edit\t\t<id> <new-name>", "Edit a todo name", 1))
	fmt.Println(printCommand("delete\t<id>", "Delete a todo", 2))
	fmt.Println(printCommand("list", "List all todos", 4))
	fmt.Println(printCommand("complete\t<id>", "Mark a todo as complete", 2))
	fmt.Println(printCommand("incomplete\t<id>", "Mark a todo as incomplete", 2))
	fmt.Println(printCommand("help", "Print usage/help", 4))
}

func printUsageWithMessage(message string) {
	fmt.Println(formatWithRed(message))
	printUsage()
	os.Exit(1)
}

func main() {
	// TODO: Load todos from JSON file if exists
	todos := Todos{}

	todos.Add("Write CLI app")
	todos.Add("Fix all GitHub issues")
	todos.Add("Go for a walk")

	if len(os.Args) < 2 {
		printUsageWithMessage("No subcommands provided")
	}

	switch os.Args[1] {
	case AddCommand:
		todos.Add(strings.Join(os.Args[2:], " "))
		renderTodosTable(todos)
	case EditCommand:
		todoId := validateArgs(os.Args, 3)
		todos.Edit(todoId, strings.Join(os.Args[3:], " "))
		renderTodosTable(todos)
	case DeleteCommand:
		todoId := validateArgs(os.Args, 3)
		todos.Delete(todoId)
		renderTodosTable(todos)
	case ListCommand:
		renderTodosTable(todos)
	case CompleteCommand:
		todoId := validateArgs(os.Args, 3)
		todos.Complete(todoId)
		renderTodosTable(todos)
	case IncompleteCommand:
		todoId := validateArgs(os.Args, 3)
		todos.Incomplete(todoId)
		renderTodosTable(todos)
	case HelpCommand:
		printUsage()
	default:
		message := fmt.Sprintf("Invalid subcommand '%v'", os.Args[1])
		printUsageWithMessage(message)
	}
}
