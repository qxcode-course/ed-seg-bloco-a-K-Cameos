package main

import (
	"bufio"
	"fmt"
	"os"
)

// cursor interativo
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		var left []rune
		var right []rune

		for _, cmd := range line {
			switch cmd {
			case 'R':
				left = append(left, '\n')
			case 'B':
				if len(left) > 0 {
					left = left[:len(left)-1]
				}
			case 'D':
				if len(right) > 0 {
					right = right[:len(right)-1]
				}
			case '<':
				if len(left) > 0 {
					right = append(right, left[len(left)-1])
					left = left[:len(left)-1]
				}
			case '>':
				if len(right) > 0 {
					left = append(left, right[len(right)-1])
					right = right[:len(right)-1]
				}
			default:
				left = append(left, cmd)
			}
		}

		fmt.Print(string(left))

		fmt.Print("|")

		for i := len(right) - 1; i >= 0; i-- {
			fmt.Print(string(right[i]))
		}
		fmt.Println()
	}
}
