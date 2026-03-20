package main
import "fmt"

type ponto struct {
    x, y int
}

func main() {
    var q int
    var d string
    fmt.Scan(&q, &d)

    cobra := make([]ponto, q)

    for i := 0; i < q; i++ {
        fmt.Scan(&cobra[i].x, &cobra[i].y)
    }

    anterior := make([]ponto, q)
    copy(anterior, cobra)

    switch d {
    case "L":
        cobra[0].x--
    case "R":
        cobra[0].x++
    case "U":
        cobra[0].y--
    case "D":
        cobra[0].y++
    }

    for i := 1; i < q; i++ {
        cobra[i] = anterior[i-1]
      }
    
    for i := 0; i < q; i++ {
        fmt.Println(cobra[i].x, cobra[i].y)
      }
}
