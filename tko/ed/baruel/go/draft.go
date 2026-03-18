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

    primeira := true

    for i := 1; i <= n; i++ {
        if freq[i] > 1 {
            for j := 0; j < freq[i]-1; j++ {
                if !primeira {
                    fmt.Print(" ")
                }
                fmt.Print(i)
                primeira = false
            }
        }
    }
    if primeira {
        fmt.Print("N")
    }
    fmt.Println()

    primeira = true

    for i := 1; i<= n; i++ {
        if freq[i] == 0 {
            if !primeira {
                fmt.Print(" ")
            }
            fmt.Print(i)
            primeira = false
        }
    }
    if primeira {
        fmt.Print("N")
    }
    fmt.Println()
}