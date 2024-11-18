package main

import (
	"fmt"
	"strings"
)

var (
	CYAN    = "\033[1;36m"
	NOCOLOR = "\033[0m"
)

func colorizeOutput(message string) string {
	return fmt.Sprintf("%v%v%v", CYAN, message, NOCOLOR)
}

func printCommand(name, description string, indentAmount int) string {
	indent := strings.Repeat("\t", indentAmount)
	return fmt.Sprintf("  %v%v%v", colorizeOutput(name), indent, description)
}
