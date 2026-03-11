package main
import "fmt"


/* recebo a lista das pessoas e quem esta com a espada
se a posicao atual for igual a posicao da espada, imprime
a pessoa com '>'. ex: 1> 2 3 4
*/

func imprimirEstado(pessoas[]int, espada int) {
    fmt.Print("[")
    for i, p := range pessoas {
        if i == espada {
            fmt.Printf(" %d>", p)
        }else {
            fmt.Printf(" %d", p)
        }
    }
    fmt.Print(" ]")
    fmt.Println()
}


func josephus(n int, e int){
    pessoas := make([]int, n)
    for i := 0; i<n; i++ {
        pessoas[i] = i+1
    }
    espada := e-1
/* cria um slice de pessoas, onde vao sendo
preenchidas a partir de 1. a posicao e ajustada 
para e-1. se comeca em 3 ajusta para 2.
*/
    for len(pessoas) > 1 {
        imprimirEstado(pessoas, espada)

        morto := (espada + 1) % len(pessoas)
// o morto eh definido por onde a espada esta + 1
// garantindo que caso seja a ultima, volta para o priemeiro indice
        pessoas = append(pessoas[:morto], pessoas[morto+1:]...)
// aqui pega tudo que esta atras do morto e junta com tudo que esta depois
// entao a posicao do morto eh cortada
        if morto < espada {
            espada --
        }
// corrije o indice, voltando a posicao da espada
        espada = (espada + 1) %len (pessoas)
    }
    imprimirEstado(pessoas, 0)
}
func main() {
    var n, e int
    fmt.Scan(&n, &e)

    josephus(n, e)

}
