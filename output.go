package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

const (
	CYAN      = "\033[1;36m"
	RED       = "\033[1;31m"
	NOCOLOR   = "\033[0m"
	CHECKMARK = "\u2714"
	CROSSMARK = "\u2718"
)

func formatWithCyan(message string) string {
	return fmt.Sprintf("%v%v%v", CYAN, message, NOCOLOR)
}

func formatWithRed(message string) string {
	return fmt.Sprintf("%v%v%v", RED, message, NOCOLOR)
}

func printCommand(name string, args []string, description string) string {
	return fmt.Sprintf("  %-28v%-17v%v", formatWithCyan(name), strings.Join(args, " "), description)
}

func printErrorMessageFatal(message string) {
	fmt.Println(formatWithRed(message))
	os.Exit(1)
}

func printTodoNotExistFatal(id int) {
	message := fmt.Sprintf("Todo with ID %v doesn't exist", id)
	printErrorMessageFatal(message)
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
			strconv.Itoa(todo.Id),
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
