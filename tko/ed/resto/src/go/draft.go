package main
import "fmt"


type Divisao struct {
    quociente int
    resto int
}

func main() {
    var numero int

    _, err := fmt.Scan(&numero)
    if err != nil {
        return
    }

    var resultados []Divisao

    for numero > 0 {
        quociente := numero / 2
        resto := numero % 2

        resultados = append(resultados, Divisao{quociente: quociente, resto: resto})

        numero = quociente
    }
    for i:= len(resultados) - 1; i >= 0; i-- {
        fmt.Println(resultados[i].quociente, resultados[i].resto)
    }    
}
