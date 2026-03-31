package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func matchingStrings(strings[]string, queries[]string) []int {
    freqMap := make(map[string]int)
    for _, str := range strings {
        freqMap[str]++
    }

    results := make([]int, len(queries))
    for i, q := range queries {
        results[i] = freqMap[q]
    }

    return results
}

    func main() {
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Split(bufio.ScanWords)

        if !scanner.Scan() {
            return
        }
        n, _ := strconv.Atoi(scanner.Text())

        stringsArr := make([]string, n)
        for i := 0; i < n; i++ {
            scanner.Scan()
            stringsArr[i] = scanner.Text()
        }

        if !scanner.Scan() {
            return
        }
        q, _ := strconv.Atoi(scanner.Text())

        queriesArr := make([]string, q)
        for i := 0; i < q; i++ {
            scanner.Scan()
            queriesArr[i] = scanner.Text()
        }

        res := matchingStrings(stringsArr, queriesArr)

        for i, val := range res {
            if i > 0 {
                fmt.Print(" ")
            }
            fmt.Print(val)
        }
        fmt.Println()
    } 
