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
	case EditCommand:
		fmt.Println("edit command")
	case DeleteCommand:
		fmt.Println("delete command")
	case ListCommand:
		renderTodosTable(todos)
	case CompleteCommand:
		fmt.Println("complete command")
	case IncompleteCommand:
		fmt.Println("incomplete command")
	case HelpCommand:
		printUsage()
	default:
		message := fmt.Sprintf("Invalid subcommand '%v'", os.Args[1])
		printUsageWithMessage(message)
	}
}
