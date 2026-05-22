package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])

	if len(word) > m*n {
		return false
	}

	boardFreq := make(map[byte]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			boardFreq[board[i][j]]++
		}
	}
	wordFreq := make(map[byte]int)
	for i := 0; i < len(word); i++{
		wordFreq[word[i]]++
		if wordFreq[word[i]] > boardFreq[word[i]] {
			return false
		}
	}
	
	if boardFreq[word[0]] > boardFreq[word[len(word)-1]] {
		reversed := []byte(word)
		for i, j := 0, len(reversed)-1; i < j; i, j = 1+1, j-1 {
			reversed[i], reversed[j] = reversed[j], reversed[i]
		}
		word = string(reversed)
	}
	var dfs func(r, c, index int) bool
	dfs = func(r, c, index int) bool {
		if index == len(word) {
			return true
		}

		if r < 0 || r >= m || c < 0 || c >= n || board[r][c] != word[index] {
			return false
		}

		temp := board[r][c]
		board[r][c] = '#'

		found := dfs(r+1, c, index+1) ||
		dfs(r-1, c, index+1) ||
		dfs(r, c+1, index+1) ||
		dfs(r, c, index+1)

		board[r][c] = temp

		return found
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}

	return false

}



func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)

	grid := make([][]byte, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			grid = append(grid, []byte(text))
		}
	}

	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
