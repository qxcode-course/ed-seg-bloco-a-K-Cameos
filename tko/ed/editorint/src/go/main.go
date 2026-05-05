package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	lines  *List[*List[rune]]
	line   *Node[*List[rune]]
	cursor *Node[rune]
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.cursor = e.line.Value.Insert(e.cursor, r)
	e.cursor = e.cursor.Next()
}

func (e *Editor) KeyLeft() {
	if e.cursor != e.line.Value.Front() { // Se o cursor não está no início da linha
		e.cursor = e.cursor.Prev() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.line != e.lines.Front() { // Se não está na primeira linha
		e.line = e.line.Prev()        // Move para a linha anterior
		e.cursor = e.line.Value.End() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyEnter() {
	newLine := NewList[rune]()

	for e.cursor != e.line.Value.End() { 
		val := e.cursor.Value  
		e.cursor = e.line.Value.Erase(e.cursor)

		newLine.PushBack(val)
	}
	e.lines.Insert(e.line.Next(), newLine)
	e.line = e.line.Next()
	e.cursor = e.line.Value.Front()
}

func (e *Editor) KeyRight() {
	if e.cursor != e.line.Value.End() {
		e.cursor = e.cursor.Next()
		return
	}

	if e.line != e.lines.Back() { 
		e.line = e.line.Next()   
		e.cursor = e.line.Value.Front() 
	}
}

func (e *Editor) KeyUp() {
	if e.line != e.lines.Front() { 
		idx := e.line.Value.IndexOf(e.cursor)

		e.line = e.line.Prev()

		e.cursor = e.line.Value.Front()
		for i := 0; i < idx && e.cursor != e.line.Value.End(); i++ {
			e.cursor = e.cursor.Next()
		}
	}
}

func (e *Editor) KeyDown() {
	if e.line != e.lines.Back() { 
		idx := e.line.Value.IndexOf(e.cursor)

		e.line = e.line.Next()

		e.cursor = e.line.Value.Front()
		for i := 0; i < idx && e.cursor != e.line.Value.End(); i++ {
			e.cursor = e.cursor.Next()
		}
	}
}

func (e *Editor) KeyBackspace() {
	if e.cursor != e.line.Value.Front() {
		prev := e.cursor.Prev()
		e.line.Value.Erase(prev)
		return
	}
	if e.line != e.lines.Front() { 
		prevLineNode := e.line.Prev()
		prevLine := prevLineNode.Value
		currentLine := e.line.Value

		var newCursor *Node[rune]

		it := currentLine.Front()
		for it != currentLine.End() {
			inserted := prevLine.Insert(prevLine.End(), it.Value)
			if newCursor == nil {
				newCursor = inserted
			}
			it = it.Next()
		}

		if newCursor == nil {
			newCursor = prevLine.End()
		}
		e.lines.Erase(e.line)
		e.line = prevLineNode
		e.cursor = newCursor
	}
}

func (e *Editor) KeyDelete() {
	if e.cursor != e.line.Value.End() {
		e.cursor = e.line.Value.Erase(e.cursor)
		return
	}
	if e.line != e.lines.Back() { 
		nextLineNode := e.line.Next()
		nextLine := nextLineNode.Value

		var firstInserted *Node[rune]

		it := nextLine.Front()
		for it != nextLine.End() {
			inserted := e.line.Value.Insert(e.line.Value.End(), it.Value)
			if firstInserted == nil {
				firstInserted = inserted
			}
			it = it.Next()
		}

		if firstInserted == nil {
			firstInserted = e.line.Value.End()
		}
		e.lines.Erase(nextLineNode)
		e.cursor = firstInserted
	}
}

func main() {
	// Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.lines = NewList[*List[rune]]()
	e.lines.PushBack(NewList[rune]())
	e.line = e.lines.Front()
	e.cursor = e.line.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.lines.Front(); line != e.lines.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '⤶'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.cursor {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
