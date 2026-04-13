package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	fila := make([]string, n)
	for i := range fila {
		fmt.Scan(&fila[i])
	}

	var m int
	fmt.Scan(&m)

	fora := make (map[string]struct{}, m)
	for i := 0; i < m;i++ {
		var num string
		fmt.Scan(&num)
		fora[num] = struct{}{}
	}

	for _, num := range fila {
		if _, ok := fora[num]; !ok {
			fmt.Print(num, " ")
		}
	}
	fmt.Println()

}
