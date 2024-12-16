package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	AddCommand        = "add"
	EditCommand       = "edit"
	DeleteCommand     = "delete"
	ListCommand       = "list"
	CompleteCommand   = "complete"
	IncompleteCommand = "incomplete"
	HelpCommand       = "help"
)

func printUsage() {
	fmt.Println("")
	fmt.Println(formatWithCyan("Usage: todo <command> [arguments]"))
	fmt.Println("\nCommands:")
	fmt.Println(printCommand("add", []string{"[name]"}, "Add a todo"))
	fmt.Println(printCommand("edit", []string{"[id]", "[new-name]"}, "Edit a todo name"))
	fmt.Println(printCommand("delete", []string{"[id]"}, "Delete a todo"))
	fmt.Println(printCommand("list", []string{}, "List all todos"))
	fmt.Println(printCommand("complete", []string{"[id]"}, "Mark a todo as complete"))
	fmt.Println(printCommand("incomplete", []string{"[id]"}, "Mark a todo as incomplete"))
	fmt.Println(printCommand("help", []string{}, "Print usage/help"))
}

func printUsageWithMessage(message string) {
	fmt.Println(formatWithRed(message))
	printUsage()
	os.Exit(1)
}

func main() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		printErrorMessageFatal("Couldn't get home directory")
	}

	todosStorePath := filepath.Join(homedir, "todo-cli", "todos.json")
	todosStoreDir := filepath.Dir(todosStorePath)

	if _, err = os.Stat(todosStorePath); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			err := os.MkdirAll(todosStoreDir, 0755)
			if err != nil {
				printErrorMessageFatal("Couldn't create todo-cli directory")
			}

			file, err := os.OpenFile(todosStorePath, os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				printErrorMessageFatal("Couldn't create todos.json")
			}
			defer file.Close()

			fmt.Printf("Created todos store at %v\n\n", todosStorePath)
		} else {
			printErrorMessageFatal(err.Error())
		}
	}

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
		todoId := validateArgs(os.Args, 3)
		todos.Edit(todoId, strings.Join(os.Args[3:], " "))
	case DeleteCommand:
		todoId := validateArgs(os.Args, 3)
		todos.Delete(todoId)
	case ListCommand:
		renderTodosTable(todos)
	case CompleteCommand:
		todoId := validateArgs(os.Args, 3)
		todos.Complete(todoId)
	case IncompleteCommand:
		todoId := validateArgs(os.Args, 3)
		todos.Incomplete(todoId)
	case HelpCommand:
		printUsage()
	default:
		message := fmt.Sprintf("Invalid subcommand '%v'", os.Args[1])
		printUsageWithMessage(message)
	}
}
