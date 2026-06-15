package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	Value int
	next *Node
	prev *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	return &LList{
		root: root,
		size: 0,
	}
}

func(l *LList) PushFront(value int) {
	n := &Node{Value: value}
	n.next = l.root.next
	n.prev = l.root
	l.root.next.prev = n
	l.root.next = n
	l.size++
}

func (l *LList) PushBack(value int) {
	n := &Node{Value: value}
	n.next = l.root
	n.prev = l.root.prev
	l.root.prev.next = n
	l.root.prev = n
	l.size++
}

func (l *LList) PopFront() {
	if l.size > 0 {
		target := l.root.next
		l.root.next = target.next
		target.next.prev = l.root
		l.size--
	}
}

func (l *LList) PopBack() {
	if l.size > 0 {
		target := l.root.prev
		l.root.prev = target.prev
		target.prev.next = l.root
		l.size--
	}
}

func (l *LList) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func (l *LList) Size() int {
	return l.size
}

func (l *LList) String() string {
	var values []string
	for n := l.root.next; n != l.root; n = n.next {
		values = append(values, strconv.Itoa(n.Value))
	}
	return "[" + strings.Join(values, ", ") + "]"
}



func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
			ll.PushBack(num)
			 }
		case "push_front":
			for _, v := range args[1:] {
			num, _ := strconv.Atoi(v)
			ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
