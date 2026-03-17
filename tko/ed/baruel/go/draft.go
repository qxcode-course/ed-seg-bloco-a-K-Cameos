package main

import (
	"fmt"
)
func main() {
    var n, m int
    fmt.Scan(&n, &m)

    freq := make([] int, n+1)
    for i := 0; i<m; i++ {
        var x int
        fmt.Scan(&x)
        freq[x]++
    }

    repetida := false
    for i := 1; i <= n; i++ {
        if freq[i] > 1 {
            if repetida {
                fmt.Println(" ")
            }
            fmt.Print(i)
            repetida = true
        }
    }
    if !repetida {
        fmt.Print("N")
    }
    fmt.Println()

    falta := false
    for i:= 1; i<=n; i++ {
        if freq[i] == 0 {
            if falta {
                fmt.Print(" ")
            }
            fmt.Print(i)
            falta = true
        }
    }
    if !falta {
        fmt.Print("N")
    }
    fmt.Println()
}
