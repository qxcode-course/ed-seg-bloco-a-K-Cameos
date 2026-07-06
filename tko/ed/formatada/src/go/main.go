package main

// está certo... porem a formatação não está igual, deve ser um problema de codificação
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// MyShow imprime a árvore binária de forma formatada.
func MyShow(node *Node, depth int) {
	indent := strings.Repeat("....", depth)

	if node == nil {
		fmt.Printf("%s#\n", indent)
		return
	}

	hasChildren := node.Left != nil || node.Right != nil

	if hasChildren {
		MyShow(node.Left, depth+1)
		fmt.Printf("%s%d\n", indent, node.Value)
		MyShow(node.Right, depth+1)
	} else {
		fmt.Printf("%s%d\n", indent, node.Value)
	}
}

// -----------------------------------------------------------------------------------
func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}
	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, history+"r")
	}
}

func create(parts *[]string) *Node {
	if len(*parts) == 0 || (*parts)[0] == "" {
		return nil
	}
	elem := (*parts)[0]
	*parts = (*parts)[1:]
	if elem == "#" {
		return nil
	}
	value, _ := strconv.Atoi(elem)
	node := &Node{Value: value}
	node.Left = create(parts)
	node.Right = create(parts)
	return node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		parts := strings.Split(text, " ")
		root := create(&parts)
		BShow(root, "") // Chama a função de impressão formatada
		MyShow(root, 0) // Chama a função de impressão personalizada
	}
}
