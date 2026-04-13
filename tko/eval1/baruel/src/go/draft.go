package main
import "fmt"

func printLista(nums []int) {
    if len(nums) == 0 {
        fmt.Println("N")
        return
    }
    for i, v := range nums {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(v)
    }
    fmt.Println()
}

func main() {
    var qt, qp int
    fmt.Scan(&qt)
    fmt.Scan(&qp)
    
    freq := make(map[int]int)

    for i := 0; i < qp; i++ {
        var y int
        fmt.Scan(&y)
        freq[y]++
    }

    var repetidas []int
    var faltando []int

    for i := 1; i <= qt; i++ {
        if freq[i] > 1 {
            for j := 0; j < freq[i]-1; j++ {
                repetidas = append(repetidas, i)
            }
        }

        if freq[i] == 0 {
            faltando = append(faltando, i)
        }
    }

    printLista(repetidas)
    printLista(faltando)
}
