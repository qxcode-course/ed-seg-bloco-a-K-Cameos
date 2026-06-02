package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" 
	"strings"
	"sort"
)

type Pair struct {
	One int
	Two int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func occurr(vet []int) []Pair {
	counts := make(map[int]int)
	for _, v := range vet {
		counts[abs(v)]++
	}

	var keys []int
	for k := range counts{
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var res []Pair
	for _, k := range keys {
		res = append(res, Pair{k, counts[k]})
	}
	return res
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return []Pair{}
	}
	var res []Pair
	currVal := vet[0]
	count := 1

	for i := 1; i < len(vet); i++ {
		if vet[i] == currVal {
			count++
		}else {
			res = append(res, Pair{currVal, count})
			currVal = vet[i]
			count = 1
		}
	}
	res = append(res, Pair{currVal, count})
	return res
}

func mnext(vet []int) []int {
	res := make([]int, len(vet))
	for i, v := range vet {
		if v > 0 {
			prevIsWoman := i > 0 && vet[i-1] < 0
			nextIsWoman := i < len(vet) - 1 && vet[i+1] < 0
			if prevIsWoman || nextIsWoman {
				res[i] = 1
			}
		}
	}
	return res
}

func alone(vet []int) []int {
	res := make([]int, len(vet))
	for i, v := range vet {
		if v > 0 {
			prevIsWoman := i > 0 && vet[i-1] < 0
			nextIsWoman := i < len(vet) - 1 && vet[i+1] < 0
			if !prevIsWoman && !nextIsWoman {
				res[i] = 1
		}
	}
}
return res
}

func couple(vet []int) int {
	counts := make(map[int]int)
	for _, v := range vet {
		counts[v]++
	}
	totalCouples := 0
	for k, countMen := range counts {
		if k > 0 {
			countWoman := counts[-k]
			if countMen < countWoman {
				totalCouples += countMen
			}else {
				totalCouples += countWoman
			}
		}
	}
	return totalCouples
}
func hasSubseq(vet []int, seq []int, pos int) bool {
	if pos+len(seq) > len(vet) {
		return false
	}
	for i := 0; i < len(seq); i++ {
		if vet[pos+i] != seq[i] {
			return false
		}
	}
	return true
}

func subseq(vet []int, seq []int) int {
	if len(seq) == 0 {
		return 0
	}
	for i := 0; i <= len(vet)-len(seq); i++ {
		if hasSubseq(vet, seq, i) {
			return i
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {
	toRemove := make(map[int]bool)
	for _, p := range posList {
		toRemove[p] = true
	}
	var res []int
	for i, v := range vet {
		if !toRemove[i] {
			res = append(res, v)
		}
	}
	return res
}

func clear(vet []int, value int) []int {
	var res []int
	for _, v := range vet {
		if v != value {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
