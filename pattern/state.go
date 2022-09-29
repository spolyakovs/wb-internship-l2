package main

import "fmt"

type state interface {
	put(*Stack, int)
	pop(*Stack) int
}

type Stack struct {
	data  []int
	state state
}

func (s *Stack) Put(elem int) {
	s.state.put(s, elem)
}

func (s *Stack) Pop() int {
	return s.state.pop(s)
}

func NewStack() *Stack {
	return &Stack{
		data:  make([]int, 0),
		state: stateEmpty{},
	}
}

type stateEmpty struct{}

func (s stateEmpty) put(stack *Stack, elem int) {
	fmt.Printf("Putting %v into empty stack\n", elem)
	stack.data = append(stack.data, elem)
	stack.state = stateNotEmpty{}
}

func (s stateEmpty) pop(stack *Stack) int {
	fmt.Println("Can't pop from empty stack")
	return 0
}

type stateNotEmpty struct{}

func (s stateNotEmpty) put(stack *Stack, elem int) {
	fmt.Printf("Putting %v into not empty stack\n", elem)
	stack.data = append(stack.data, elem)
	stack.state = stateNotEmpty{}
}

func (s stateNotEmpty) pop(stack *Stack) int {
	result := stack.data[len(stack.data)-1]
	fmt.Printf("Getting %v from not empty stack\n", result)
	stack.data = stack.data[:len(stack.data)-1]
	if len(stack.data) == 0 {
		stack.state = stateEmpty{}
	}
	return result
}

func StateExample() {
	st := NewStack()

	st.Pop()
	st.Put(1)
	st.Put(2)
	st.Pop()
}
