package main

import (
	"fmt"
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

func printCommand(name string, args []string, description string) string {
	return fmt.Sprintf("  %-23v%-17v%v", formatWithCyan(name), strings.Join(args, " "), description)
}

func printErrorMessageFatal(message string) {
	fmt.Println(formatWithRed(message))
	os.Exit(1)
}

func renderTodosTable(todos Todos) {
	table := tablewriter.NewWriter(os.Stdout)
	data := [][]string{}

	for _, todo := range todos {
		var completedAt string
		var completedStatus string

		if todo.CompletedAt == nil {
			completedAt = ""
		} else {
			completedAt = todo.CompletedAt.Format(time.Stamp)
		}

		if todo.Completed == true {
			completedStatus = CHECKMARK
		} else {
			completedStatus = CROSSMARK
		}

		data = append(data, []string{
			strconv.Itoa(todo.Id - 1),
			todo.Name,
			completedStatus,
			completedAt,
			todo.CreatedAt.Format(time.Stamp),
		})
	}

	table.SetHeader([]string{"ID", "Name", "Completed", "Completed at", "Created at"})
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(data)
	table.Render()
}
