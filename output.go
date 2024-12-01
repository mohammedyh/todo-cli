package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func formatWithCyan(message string) string {
	return fmt.Sprintf("%v%v%v", CYAN, message, NOCOLOR)
}

func formatWithRed(message string) string {
	return fmt.Sprintf("\n%v%v%v", RED, message, NOCOLOR)
}

func printCommand(name, description string, indentAmount int) string {
	indent := strings.Repeat("\t", indentAmount)
	return fmt.Sprintf("  %v%v%v", formatWithCyan(name), indent, description)
}

func printSliceToJSON(slice Todos) {
	sliceJSON, err := json.MarshalIndent(slice, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(sliceJSON))
}
