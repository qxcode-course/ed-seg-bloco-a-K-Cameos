package main
import "fmt"
func main() {
    var s string
    var L int

    if _, err := fmt.Scan(&s, &L); err != nil {
        return
    }

    T := L + 1

    tamp := make([]byte, T)
    for i := 0; i < T; i++ {
        tamp[i] = '.'
    }

    jap := make([]bool, L+1)

    for i := 0; i < len(s); i++ {
        if s[i] != '.' {
            tamp[i%T] = s[i]

            digito := int (s[i] - '0')
            if digito <= L {
                jap[digito] = true
            }
        }
    }
    var njap []byte
    for d := 0; d <= L; d++ {
        if !jap[d] {
            njap = append(njap, byte(d+'0'))
        }
    }
    uIdx := 0
    for i := 0; i < T; i++ {
        if tamp[i] == '.' {
            if uIdx < len(njap) {
                tamp[i] = njap[uIdx]
                uIdx++
            }
        }
    }
    res := make([]byte, len(s))
    for i := 0; i< len(s); i++ {
        res[i] = tamp[i%T]
    }

    fmt.Println(string(res))
}