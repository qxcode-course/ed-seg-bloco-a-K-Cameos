
package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    animais := make([] int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&animais[i])
    }
    
    ncasais := make(map[int]int)
    pares := 0

    for _, animal := range animais{
        oposto := -animal

        if ncasais[oposto] > 0{
            ncasais[oposto]--
            pares++
        }else {
            ncasais[animal]++
        }
    }

    fmt.Println(pares)
}



