package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}}
}

func (s *Stack[T]) Top() T {
	if len(s.data) == 0 {
		panic("empty stack")
	}
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("empty")
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value
}

func precedencia(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	}
	return 0
}

func infixToPostfix(expression string) string {
	tokens := strings.Fields(expression)
	stack := NewStack[string]()
	var output []string

	for _, token := range tokens {
		prec := precedencia(token)

		if prec > 0 {
			for !stack.IsEmpty() && precedencia(stack.Top()) >= prec {
				output = append(output, stack.Pop())
			}
			stack.Push(token)
		} else {
			output = append(output, token)
		}
	}
	for !stack.IsEmpty() {
		output = append(output, stack.Pop())
	}

	return strings.Join(output, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		resultado := infixToPostfix(line)
		fmt.Println(resultado)
	}
}
