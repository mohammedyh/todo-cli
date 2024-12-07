package main

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

type Todo struct {
	Id          int
	Name        string
	Completed   bool
	CompletedAt *time.Time
	CreatedAt   time.Time
}

type Todos []Todo

func (todos *Todos) Add(name string) {
	*todos = append(*todos, Todo{
		Id:          len(*todos) + 1,
		Name:        name,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	})
}

func (todos *Todos) Edit(id int, name string) {
	if id > len(*todos)-1 {
		message := fmt.Sprintf("Todo with ID %v doesn't exist", id)
		printErrorMessageFatal(message)
	}

	if len(strings.TrimSpace(name)) == 0 {
		printErrorMessageFatal("New todo name is invalid or empty")
	}

	(*todos)[id].Name = name
}

func (todos *Todos) Complete(id int) {
	if id > len(*todos)-1 {
		message := fmt.Sprintf("Todo with ID %v doesn't exist", id)
		printErrorMessageFatal(message)
	}

	if (*todos)[id].Completed == true {
		fmt.Println(formatWithRed("Todo already marked as complete"))
	}

	now := time.Now()

	(*todos)[id].Completed = true
	(*todos)[id].CompletedAt = &now
}

func (todos *Todos) Incomplete(id int) {
	if id > len(*todos)-1 {
		message := fmt.Sprintf("Todo with ID %v doesn't exist", id)
		printErrorMessageFatal(message)
	}

	if (*todos)[id].Completed == false {
		fmt.Println(formatWithRed("Todo already marked as incomplete"))
	}

	(*todos)[id].Completed = false
	(*todos)[id].CompletedAt = nil
}

func (todos *Todos) Delete(index int) {
	if index > len(*todos)-1 || index < 0 {
		printErrorMessageFatal("index out of range")
	}

	*todos = slices.Concat((*todos)[:index], (*todos)[index+1:])
}
