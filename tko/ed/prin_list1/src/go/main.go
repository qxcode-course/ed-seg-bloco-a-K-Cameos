package main

import (
	"fmt"
	"strings"
	
)
type DNode[T comparable] struct {
	Value				T
	next, prev, root *DNode[T]
}

func (n *DNode[T]) Next() *DNode[T] {
	if n == n.root {
		return n
	}
	return n.next
}

func (n *DNode[T]) Prev() *DNode[T] {
	if n == n.root {
		return n
	}
	return n.prev
}

type DList[T comparable] struct {
	root *DNode[T]
	size int
}

func NewDList[T comparable]() *DList[T] {
	root := &DNode[T]{}
	root.next = root
	root.prev = root
	root.root = root
	return &DList[T]{root: root, size: 0}
}

func (l *DList[T]) PushBack(value T) {
	l.Insert(l.root, value)
}

func  (l *DList[T]) PopBack() {
	if l.size == 0 {
		return
	}
	l.Erase(l.root.prev)
}

func (l *DList[T]) PopFront() {
	if l.size == 0 {
		return
	}
	l.Erase(l.root.next)
}

func (l *DList[T]) PushFront(value T) {
	l.Insert(l.root.next, value)
}

func (l *DList[T]) Insert(it *DNode[T], value T) *DNode[T] {
	n := &DNode[T]{Value: value, root: l.root}
	n.prev = it.prev
	n.next = it
	it.prev.next = n
	it.prev = n
	l.size++
	return n
}

func (l *DList[T]) Erase(it *DNode[T]) {
	if it == l.root || it == nil {
		return 
	}
	it.prev.next = it.next
	it.next.prev = it.prev
	l.size--
}

func (l *DList[T]) String() string {
	values := []string{}
	for n := l.root.next; n != l.root; n = n.next {
		values = append(values, fmt.Sprint(n.Value))
	}
	return "[" + strings.Join(values, ", ") + "]"
}

func (l *DList[T]) Size() int {
	return l.size
}

func (l *DList[T]) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func (l *DList[T]) Front() *DNode[T] {
	return l.root.next
}

func (l *DList[T]) Back() *DNode[T] {
	return l.root.prev
}

func (l DList[T]) End() *DNode[T] {
	return l.root
}


// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	var parts []string
	for n := l.Front(); n != l.End(); n = n.next {
		if n == sword {
			parts = append(parts, fmt.Sprintf("%d>", n.Value))
		}else {
			parts = append(parts, fmt.Sprintf("%d", n.Value))
		}
	}
	return "[ " + strings.Join(parts, " ") + " ]"
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it.next == l.root {
		return l.root.next
	}
	return it.next
}

func main() {
	var qtd, chosen int
	if _, err := fmt.Scan(&qtd, &chosen); err != nil {
		return
	}
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	
	for i := 0; i < chosen-1; i++ {
		sword = Next(l, sword)
	}

	for l.Size() > 1 {
		fmt.Println(ToStr(l, sword))
	}

	for l.Size() > 1 {
		fmt.Println(ToStr(l,sword))

		toKill := Next(l, sword)
		l.Erase(toKill)
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
