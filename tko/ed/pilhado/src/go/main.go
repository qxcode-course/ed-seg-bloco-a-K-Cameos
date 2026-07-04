package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}}
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *Stack[T]) Top() T {
	if len(s.data) == 0 {
		panic("stack is empty")
	}
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value
}

type Pos struct {
	l, c int
}

func escapeMaze(grid [][]rune, start, end Pos) {
	caminho := NewStack[Pos]()
	caminho.Push(start)

	dirs := []Pos{{-1, 0}, {1,0}, {0, -1}, {0, 1}}

	for !caminho.IsEmpty() {
		atual := caminho.Top()

		grid[atual.l][atual.c] = '.'

		if atual.l == end.l && atual.c == end.c {
			break
		}

		achouVizinho := false 
		for _, d := range dirs {
			nl, nc := atual.l+d.l, atual.c+d.c

			if nl >= 0 && nl < len(grid) && nc >= 0 && nc < len(grid[0]) {
				if grid[nl][nc] == ' ' {
					caminho.Push(Pos{nl, nc})
					achouVizinho = true
					break
				}
			}
		}

		if !achouVizinho {
			grid[atual.l][atual.c] = 'b'
			caminho.Pop()
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'b' {
				grid[i][j] = ' '
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}

	var nl, nc int
	fmt.Sscanf(scanner.Text(), "%d %d", &nl, &nc)

	grid := make([][]rune, nl)
	var start, end Pos

	for i := 0; i < nl; i++ {
		scanner.Scan()
		line := []rune(scanner.Text())

		if len(line) < nc {
			padded := make([]rune, nc)
			copy(padded, line)
			for j := len(line); j < nc; j++ {
				padded[j] = ' '
			}
			line = padded
		}

		grid[i] = line

		for j := 0; j < len(line); j++ {
			if line[j] == 'I' {
				start = Pos{i, j}
				grid[i][j] = ' '
			} else if line[j] == 'F' {
				end = Pos{i, j}
				grid[i][j] = ' '
			}
		}
	}

	escapeMaze(grid, start, end)

	for i := 0; i < nl; i++ {
		fmt.Println(string(grid[i]))
	}
}
