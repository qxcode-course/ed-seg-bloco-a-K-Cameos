package main
import "fmt"
func contandoPares(animais []int) int {
    ncasados := make(map[int]int)
    totalCasais := 0

    for _, animal := range animais {
        oposto := -animal

        if ncasados[oposto] > 0 {
            ncasados[oposto]--
            totalCasais++
        } else {
            ncasados[animal]++
        }
    }

    return totalCasais
}

func lendoEntrada() []int {
    var n int
    fmt.Scan(&n)

    animais := make([]int, n)
    for i := range animais {
        fmt.Scan(&animais[i])
    }

    return animais
}

func main() {
    animais := lendoEntrada()
    resultado := contandoPares(animais)
    fmt.Println(resultado)    
}
