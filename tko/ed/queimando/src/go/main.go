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


func burnTrees(grid [][]rune, l, c int) {
	nl := len(grid)
	if nl == 0 {
		return
	}
	nc := len(grid[0])

	stack := NewStack[Pos]()
	
	stack.Push(Pos{l, c})

	for !stack.IsEmpty() {
		curr := stack.Pop()

		if curr.l >= 0 && curr.l < nl && curr.c >= 0 && curr.c < nc && grid[curr.l][curr.c] == '#' {
			grid[curr.l][curr.c] = 'o'
			
			stack.Push(Pos{curr.l - 1, curr.c})
			stack.Push(Pos{curr.l + 1, curr.c})
			stack.Push(Pos{curr.l, curr.c - 1})
			stack.Push(Pos{curr.l, curr.c + 1})

		}
	}

	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
