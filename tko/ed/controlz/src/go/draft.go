package main
import (
    "bufio"
    "fmt"
    "os"
    
)

type State struct {
    left []rune
    right []rune
}

type Editor struct {
    left []rune
    right []rune
    undo []State
    redo []State
}

func NewEditor() *Editor {
    return &Editor{}
}

func clone(r []rune) []rune {
    if len(r) == 0 {
        return nil
    }
    res := make([]rune, len(r))
    copy(res, r)
    return res
}

func (e *Editor) saveState() {
    e.undo = append(e.undo, State{clone(e.left), clone(e.right)})
    e.redo = nil
}

func (e *Editor) Process(cmd rune) {
    switch  {
    case (cmd >= 'a' && cmd <= 'z') || cmd == '-':
    e.saveState()
    e.left = append(e.left, cmd)
    case cmd == 'R':
        e.saveState()
        e.left = append(e.left, '\n')
    case cmd == 'B':
        if len(e.left) > 0 {
            e.saveState()
            e.left = e.left[:len(e.left)-1]
        }
    case cmd == 'D':
        if len(e.right) > 0 {
            e.saveState()
            e.right = e.right[:len(e.right)-1]
        }
    case cmd == '<':
        if len(e.left) > 0 {
            e.right = append(e.right, e.left[len(e.left)-1])
            e.left = e.left[:len(e.left)-1]
        }
    case cmd == '>':
        if len(e.right) > 0 {
            e.left = append(e.left, e.right[len(e.right)-1])
            e.right = e.right[:len(e.right)-1]
        }
    case cmd == 'Z':
        if len(e.undo) > 0 {
            e.redo = append(e.redo, State{clone(e.left), clone(e.right)})

            last := e.undo[len(e.undo)-1]
            e.undo = e.undo[:len(e.undo)-1]

            e.left, e.right = clone(last.left), clone(last.right)
        }
    case cmd == 'Y':
        if len(e.redo) > 0 {
            e.undo = append(e.undo, State{clone(e.left), clone(e.right)})

            last := e.redo[len(e.redo)-1]
            e.redo = e.redo[:len(e.redo)-1]
            e.left, e.right = clone(last.left), clone(last.right)
        }
    }
}

func (e *Editor) Print() {
    fmt.Print(string(e.left))
    fmt.Print("|")

    for i := len(e.right) - 1; i >= 0; i-- {
        fmt.Print(string(e.right[i]))
    }
    fmt.Println()
}


func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        line := scanner.Text()
        if len(line) == 0 {
            continue
        }

        ed := NewEditor()
        for _, ch := range line {
            ed.Process(ch)
        }
        ed.Print()
    }
}