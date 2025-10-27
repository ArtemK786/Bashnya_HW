package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) PushBack(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) PushFront(i int) {
	new_stack := Stack{}
	new_stack.items = append(new_stack.items, i)
	for i = 0; i != len(s.items); i++ {
		new_stack.items = append(new_stack.items, s.items[i])
	}
	s.items = new_stack.items

}

func (s *Stack) PopFront() {
	s.items = s.items[1:]
}

func (s *Stack) PopBack() {
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) IsEmpty() string {
	lenStack := len(s.items)
	if lenStack != 0 {
		return "Not empty"
	} else {
		return "EMPTY"
	}
}

func (s *Stack) Size() int {
	lenStack := len(s.items)
	return lenStack
}

func (s *Stack) Clear() {
	for len(s.items) != 0 {
		s.PopBack()
	}
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.PushBack(2)
	myStack.PushBack(3)
	myStack.PushFront(100)
	fmt.Println(myStack)
	myStack.PopFront()
	fmt.Println(myStack)
	myStack.PopBack()
	fmt.Println(myStack)
	fmt.Println(myStack.IsEmpty())
	myStack.Clear()
	fmt.Println(myStack)
}
