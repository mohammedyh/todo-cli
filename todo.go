package main

import (
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
