package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"time"
)

// change Id to type uint and convert where needed
type Todo struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Completed   bool       `json:"completed"`
	CompletedAt *time.Time `json:"completedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
}

type Todos []Todo

func (todos *Todos) Add(name string) {
	validateTodoName(name)

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

	validateTodoName(name)
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

func (todos *Todos) Load() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		printErrorMessageFatal("Couldn't get home directory")
	}

	todosStorePath := filepath.Join(homedir, "todo-cli", "todos.json")
	data, err := os.ReadFile(todosStorePath)
	if err != nil {
		printErrorMessageFatal(err.Error())
	}

	err = json.Unmarshal(data, todos)
	if err != nil {
		printErrorMessageFatal(err.Error())
	}
}

func (todos *Todos) Save() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		printErrorMessageFatal("Couldn't get home directory")
	}

	todosStorePath := filepath.Join(homedir, "todo-cli", "todos.json")
	todosJson, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		printErrorMessageFatal(err.Error())
	}

	err = os.WriteFile(todosStorePath, todosJson, os.FileMode(os.O_RDWR))
	if err != nil {
		printErrorMessageFatal(err.Error())
	}
}
