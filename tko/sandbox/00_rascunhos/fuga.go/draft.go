package main
import "fmt"
func main() {
    var h, p, f, d int
    fmt.Scan(&h, &p, &f, &d)

    atual := f
    fugiu := false

    for {
        atual += d

        if atual > 15 {
            atual = 0
        }else if atual < 0 {
            atual = 15
        }
        
        if atual == h {
            fugiu = true
            break
        }else if atual == p {
            fugiu = false
            break
        }
    }
    if fugiu {
        fmt.Println("S")
    }else {
        fmt.Println("N")
    }

}
