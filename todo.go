package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"time"
)

type Todo struct {
	Id          uint       `json:"id"`
	Name        string     `json:"name"`
	Completed   bool       `json:"completed"`
	CompletedAt *time.Time `json:"completedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
}

type Todos []Todo

var nextId uint

func (todos *Todos) Add(name string) {
	validateTodoName(name)

	if len(*todos) == 0 {
		nextId = 0
	} else {
		nextId = (*todos)[len(*todos)-1].Id + 1
	}

	*todos = append(*todos, Todo{
		Id:          nextId,
		Name:        name,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	})
}

func (todos *Todos) Edit(id uint, name string) {
	todoIndex := todos.findById(id)
	if todoIndex == -1 {
		printTodoNotExistFatal(id)
	}

	validateTodoName(name)
	(*todos)[todoIndex].Name = name
}

func (todos *Todos) Complete(id uint) {
	todoIndex := todos.findById(id)
	if todoIndex == -1 {
		printTodoNotExistFatal(id)
	}

	if (*todos)[todoIndex].Completed == true {
		fmt.Println(formatWithRed("Todo already marked as complete"))
	}

	now := time.Now()

	(*todos)[todoIndex].Completed = true
	(*todos)[todoIndex].CompletedAt = &now
}

func (todos *Todos) Incomplete(id uint) {
	todoIndex := todos.findById(id)
	if todoIndex == -1 {
		printTodoNotExistFatal(id)
	}

	if (*todos)[todoIndex].Completed == false {
		fmt.Println(formatWithRed("Todo already marked as incomplete"))
	}

	(*todos)[todoIndex].Completed = false
	(*todos)[todoIndex].CompletedAt = nil
}

func (todos *Todos) ClearCompleted() {
	var completedCount int
	for _, todo := range *todos {
		if todo.Completed {
			completedCount++
			todos.Delete(todo.Id)
		}
	}

	if completedCount == 0 {
		fmt.Println(formatWithRed("No more completed todos to clear"))
	}
}

func (todos *Todos) ClearAll() {
	if len(*todos) == 0 {
		printErrorMessageFatal("No todos to clear")
	}
	*todos = (*todos)[:0]
}

func (todos *Todos) Delete(id uint) {
	todoIndex := todos.findById(id)
	if todoIndex == -1 {
		printTodoNotExistFatal(id)
	}

	*todos = slices.Concat((*todos)[:todoIndex], (*todos)[todoIndex+1:])
}

func (todos *Todos) Load() error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	todosStorePath := filepath.Join(homedir, "todo-cli", "todos.json")
	data, err := os.ReadFile(todosStorePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, todos)
	if err != nil {
		return err
	}
	return nil
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

func (todos *Todos) findById(id uint) int {
	return slices.IndexFunc(*todos, func(todo Todo) bool {
		return todo.Id == id
	})
}
