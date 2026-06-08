package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Set struct {
	data []int
	size int
	capacity int
}

func NewSet(capacity int) *Set {
	if capacity <= 0 {
		capacity = 1
	}
	return &Set{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (s *Set) reserve(newCapacity int) {
	if newCapacity > s.capacity {
		newData := make ([]int, newCapacity)
		copy(newData, s.data[:s.size])
		s.data = newData
		s.capacity = newCapacity
	}
}

func (s *Set) binaySearch(value int) int{
	low, high := 0, s.size-1
	for low <= high {
		mid := low + (high-low)/2
		if s.data[mid] == value {
			return mid
		}else if s.data[mid] < value {
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	return -1
}

func (s *Set) findInsertIndex(value int) int {
	low, high := 0, s.size-1
	for low <= high {
		mid := low + (high-low)/2
		if s.data[mid] == value {
			return mid
		}else if s.data[mid] < value {
			low = mid + 1
		}else {
			high = mid - 1
		}
	}
	return low
}

func (s *Set) insert(value int, index int) error {
	if s.size == s.capacity {
		s.reserve(s.capacity * 2)
	}
	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[index] = value
	s.size++
	return nil
}

func (s *Set) erase(index int) error {
	for i := index; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.size--
	return nil
}

func (s *Set) Insert(value int) {
	idx := s.findInsertIndex(value)
	if idx < s.size && s.data[idx] == value {
		return
	}
	s.insert(value, idx)
}

func (s *Set) Contains(value int) bool {
	return s.binaySearch(value) != -1
}

func (s *Set) Erase(value int) bool {
	idx := s.binaySearch(value)
	if idx == -1 {
		return false
	}
	s.erase(idx)
	return true
}

func (s *Set) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < s.size; i++ {
		sb.WriteString(strconv.Itoa(s.data[i]))
		if i < s.size-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(1)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
			value, _ := strconv.Atoi(part)
			v.Insert(value)
			}
		case "show":
			fmt.Println(v.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if !v.Erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			}else {
				fmt.Println("false")
			}
		case "clear":
			v = NewSet(v.capacity)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
