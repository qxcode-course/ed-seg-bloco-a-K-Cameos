package main

import (
	"bufio"
	"fmt"
	"os"
)

func burnTrees(grid [][]rune, l, c int) {
	_, _, _ = grid, l, c
	nl := len(grid)
	nc := len(grid[0])
	_, _ = nl, nc
	if c < 0 || c >= nc || l < 0 || l >= nl {
		return
	}
	if grid[l][c] != '#' {
		return
	}
	grid[l][c] = 'o'
	burnTrees(grid, l-1, c)
	burnTrees(grid, l+1, c)
	burnTrees(grid, l, c-1)
	burnTrees(grid, l, c+1)
	
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

func showGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
