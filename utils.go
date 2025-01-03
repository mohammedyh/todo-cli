package main

import (
	"strconv"
	"strings"
)

func validateArgs(args []string, requiredArgsLength int) uint {
	if len(args) < requiredArgsLength {
		printErrorMessageFatal("Not enough arguments provided")
	}

	todoId, err := strconv.Atoi(args[2])

	if err != nil {
		printErrorMessageFatal("Invalid Todo ID")
	}

	return uint(todoId)
}

func validateTodoName(name string) {
	if len(strings.TrimSpace(name)) == 0 {
		printErrorMessageFatal("Todo name cannot be empty")
	}
}
