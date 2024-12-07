package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printUsage() {
	fmt.Println("")
	fmt.Println(formatWithCyan("Usage: todo <command> [arguments]"))
	fmt.Println("\nCommands:")
	fmt.Println(printCommand("add", "Add a todo", 2))
	fmt.Println(printCommand("edit", "Edit a todo", 2))
	fmt.Println(printCommand("delete", "Delete a todo", 1))
	fmt.Println(printCommand("list", "List all todos", 2))
	fmt.Println(printCommand("complete", "Mark a todo as complete", 1))
	fmt.Println(printCommand("incomplete", "Mark a todo as incomplete", 1))
	fmt.Println(printCommand("help", "Print usage/help", 2))
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
		fmt.Println(todos)
	case EditCommand:
		if len(os.Args) < 3 {
			printErrorMessageFatal("Todo ID not provided")
		}

		todoId, err := strconv.Atoi(os.Args[2])

		if err != nil {
			printErrorMessageFatal("Invalid Todo ID")
		}

		todos.Edit(todoId, strings.Join(os.Args[3:], " "))
	case DeleteCommand:
		if len(os.Args) < 3 {
			printErrorMessageFatal("Todo ID not provided")
		}

		todoId, err := strconv.Atoi(os.Args[2])

		if err != nil {
			printErrorMessageFatal("Invalid Todo ID")
		}

		todos.Delete(todoId)
	case ListCommand:
		renderTodosTable(todos)
	case CompleteCommand:
		todoId, err := strconv.Atoi(os.Args[2])

		if err != nil {
			printErrorMessageFatal("Invalid Todo ID")
		}

		todos.Complete(todoId)
	case IncompleteCommand:
		todoId, err := strconv.Atoi(os.Args[2])

		if err != nil {
			printErrorMessageFatal("Invalid Todo ID")
		}

		todos.Incomplete(todoId)
	case HelpCommand:
		printUsage()
	default:
		message := fmt.Sprintf("Invalid subcommand '%v'", os.Args[1])
		printUsageWithMessage(message)
	}
}
