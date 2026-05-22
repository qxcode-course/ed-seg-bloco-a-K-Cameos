package main
import "fmt"

func eh_primo(num int, divisor int) bool {
    if num < 2 {
        return false
    }
    if divisor*divisor > num {
        return true
    }
    if num%divisor == 0 {
        return false
    }
    return eh_primo(num, divisor+1)
}

func gprimos(n int, atual int, primos []int) []int {
    if len(primos) == n {
        return primos
    }
    if eh_primo(atual, 2) {
        primos = append(primos, atual)
    }
    return gprimos(n, atual+1, primos)
}

func main() {
    var n int
    _, err := fmt.Scan(&n)
    if err != nil {
        fmt.Println("[]")
        return
    }

    resultado := gprimos(n, 2, []int{})
    fmt.Printf("[")
    for i, p := range resultado {
        if i > 0 {
            fmt.Printf(", ")
        }
        fmt.Print(p)
    }
    fmt.Printf("]\n")
}
