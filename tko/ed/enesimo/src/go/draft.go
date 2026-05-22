package main
import "fmt"
func aux_primo(num, divisor int) bool {
    if num < 2 {
        return false
    }
    if divisor*divisor > num {
        return true
    }
    if num%divisor == 0{
        return false
    }
    return aux_primo(num, divisor+1)
}

func eh_primo(num int) bool {
    return aux_primo(num, 2)
}

func aux_ene(n, atual, encontrados int) int {
    if eh_primo(atual) {
        encontrados++
        if encontrados == n {
            return atual
        }
    }
    return aux_ene(n, atual+1, encontrados)
}
func n_primo(n int) int {
    return aux_ene(n, 2, 0)
}
func main() {
    var n int 

    _, err := fmt.Scan(&n)
    if err == nil && n > 0 {
        fmt.Println(n_primo(n))
    } 
}
