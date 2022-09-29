package main

import "fmt"

type Command interface {
	Execute()
}

type action struct {
	f func()
}

func NewAction(f func()) *action {
	return &action{f: f}
}

func (a *action) Execute() {
	a.f()
}

func CommandExample() {
	a := NewAction(func() {
		fmt.Println("Performing some action")
	})

	a.Execute()
}
