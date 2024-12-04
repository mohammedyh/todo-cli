package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

func formatWithCyan(message string) string {
	return fmt.Sprintf("%v%v%v", CYAN, message, NOCOLOR)
}

func formatWithRed(message string) string {
	return fmt.Sprintf("%v%v%v", RED, message, NOCOLOR)
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

func renderTodosTable(todos Todos) {
	table := tablewriter.NewWriter(os.Stdout)
	data := [][]string{}

	for _, todo := range todos {
		var completedAt string

		if todo.CompletedAt == nil {
			completedAt = "Not done"
		}

		data = append(data, []string{
			strconv.Itoa(todo.Id),
			todo.Name,
			strconv.FormatBool(todo.Completed),
			completedAt,
			todo.CreatedAt.Format(time.Stamp),
		})
	}

	table.SetHeader([]string{"ID", "Name", "Completed", "Completed at", "Created at"})
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(data)
	table.Render()
}
