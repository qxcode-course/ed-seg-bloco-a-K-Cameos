package main
import "fmt"
func main() {
    var nome string
    var idade int
    fmt.Scanf("%s\n%d\n", &nome, &idade)

    if idade < 12 {
        fmt.Printf("%s eh crianca\n", nome)
    }else if idade < 18 {
        fmt.Printf("%s eh jovem\n", nome)
    }else if idade < 65 {
        fmt.Printf("%s eh adulto\n", nome)
    }else if idade < 1000 {
        fmt.Printf("%s eh idoso\n", nome)
    }else {
        fmt.Printf("%s eh mumia\n", nome)
    }
}
