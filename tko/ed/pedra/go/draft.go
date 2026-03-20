package main

import ("fmt" 
"math")
func main() {
    var n int
    fmt.Scan(&n)
    
    maiorPontuacao := 101
    indiceVencedor := -1

    for i := 0; i < n; i++{
        var a, b int
        fmt.Scan(&a, &b)

        if a < 10 || b < 10 {
            continue
        }

        diferenca := int(math.Abs(float64(a-b)))

        if diferenca < maiorPontuacao {
            maiorPontuacao = diferenca
            indiceVencedor = i
        }
    }
    if indiceVencedor == -1 {
        fmt.Println("sem ganhador")
    }else {
        fmt.Println(indiceVencedor)
    }


}
