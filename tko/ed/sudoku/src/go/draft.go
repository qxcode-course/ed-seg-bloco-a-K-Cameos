package main
import (
    "fmt"
    "math"
)

func isValid(board [][]byte, row, col int, num byte, n int) bool {
    for i := 0; i < n; i++ {
        if board[row][i] == num {
            return false
        }
    }

    for i := 0; i < n; i++ {
        if board[i][col] == num {
            return false
        }
    }

    subSize := int(math.Sqrt(float64(n)))
    startRow := (row / subSize) * subSize
    startCol := (col / subSize) * subSize

    for i := 0; i < subSize; i++ {
        for j := 0; j < subSize; j++ {
            if board[startRow+i][startCol+j] == num {
                return false
            }
        }
    }
    return true
}

func solve(board [][]byte, index, n int) bool {
    if index == n*n {
        return true
    }

    row := index / n 
    col := index % n 

    if board[row][col] != '.' {
        return solve(board, index+1, n)
    }

    for i := 1; i<= n; i++ {
        num := byte(i+'0')
        if isValid(board, row, col, num, n) {
            board[row][col] = num
            
            if solve(board, index+1, n) {
                return true
            }

            board[row][col] = '.'
        }
    }

    return false
}


func main() {
    var n int

    if _, err := fmt.Scan(&n); err != nil {
        return
    }

    board := make([][]byte, n)
    for i := 0; i < n; i++ {
        var rowStr string
        fmt.Scan(&rowStr)
        board[i] = []byte(rowStrv) 
    }

    solve(board, 0, n)

    for i := 0; i < n; i++ {
        fmt.Println(string(board[i]))
    }
}