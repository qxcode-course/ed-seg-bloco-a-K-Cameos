package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	
	if !scanner.Scan() {
		return
	}
	n, _ := strconv.Atoi(scanner.Text())
	
	filaInicial := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		filaInicial[i] = scanner.Text()
	}
	
	if !scanner.Scan() {
		return
	}
	m, _ := strconv.Atoi(scanner.Text())
	
	saiuMap := make(map[string]bool)
	for i := 0; i < m; i++ {
		scanner.Scan()
		saiuMap[scanner.Text()] = true
	}
	
	for _, id := range filaInicial {
		if !saiuMap[id] {
			fmt.Printf("%s ", id)
		}
	}
	
	fmt.Println()
}
