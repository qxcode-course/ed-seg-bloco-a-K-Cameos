package main

import (
	"container/list"
	"fmt"
	"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *list.List, sword *list.Element) string {
	var parts []string

	for e := l.Front(); e != nil; e = e.Next() {
		val := e.Value.(int)

		if e == sword {
			if val > 0 {
				parts = append(parts, fmt.Sprintf("%d>", val))
			} else {
				parts = append(parts, fmt.Sprintf("<%d", val))
			}

		} else {
			parts = append(parts, fmt.Sprintf("%d", val))
		}
	}

	return "[ " + strings.Join(parts, " ") + " ]"
}

// move para frente na lista circular
func Next(l *list.List, it *list.Element) *list.Element {
	if it.Next() != nil {
		return it.Next()
	}
	return l.Front()
}

// move para tras na lista circular
func Prev(l *list.List, it *list.Element) *list.Element {
	if it.Prev() != nil {
		return it.Prev()
	}
	return l.Back()
}

func main() {
	var qtd, chosen, fase int
	if _, err := fmt.Scan(&qtd, &chosen, &fase); err != nil {
		return
	}
	l := list.New()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i * fase)
		fase = -fase
	}
	sword := l.Front()
	for i := 0; i < chosen-1; i++ {
		sword = Next(l, sword)
	}

	for l.Len() > 1 {
		fmt.Println(ToStr(l, sword))

		val := sword.Value.(int)

		if val > 0 {
			l.Remove(Next(l, sword))
			sword = Next(l, sword)
		} else {
			l.Remove(Prev(l, sword))
			sword = Prev(l, sword)
		}
	}
	fmt.Println(ToStr(l, sword))
}
