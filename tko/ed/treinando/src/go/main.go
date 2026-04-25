package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}

	var rec func(v []int) string 
	rec = func(v []int) string {
		if len(v) == 1 {
			return strconv.Itoa(v[0])
		}
		return strconv.Itoa(v[0]) + ", " + rec(v[1:])
	}

	return "[" + rec(vet) + "]"
}

func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
//
	var rec func(v []int) string
	rec = func(v []int) string {
		if len(v) == 1 {
			return strconv.Itoa(v[0])
		}
		return rec(v[1:]) + ", " + strconv.Itoa(v[0])
	}
	return "[" + rec(vet) + "]"
}
// reverse: inverte os elementos do slice
func reverse(vet []int) {
	if len(vet) <= 1 {
		return
	}

	vet[0], vet[len(vet)-1] = vet[len(vet)-1], vet[0]
	reverse(vet[1 : len(vet)-1])
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	return vet[0] + sum(vet[1:])
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0 {
		return 1
	}
	return vet[0] * mult(vet[1:])
}

// min: retorna o índice e valor do menor valor
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}
	var rec func(v []int) (int, int)
	rec = func(v []int) (int, int) {
		if len(v) == 1 {
			return 0, v[0]
		}

		idxRest, valRest := rec(v[1:])
		if v[0] <= valRest {
			return 0, v[0]
		}
		return idxRest + 1, valRest
	}
	idx, _ := rec(vet)
	return idx
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
