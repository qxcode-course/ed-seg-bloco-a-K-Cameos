package main
import "fmt"

func solve(row, cols, diag1, diag2, n int) int {
    if row == n {
        return 1
    }

    count := 0

    disp := ((1 << n) - 1) &^ (cols | diag1 | diag2)

    for disp != 0 {
        pos := disp & -disp

        disp &= disp - 1

        count += solve(
            row+1,
            cols|pos,
            (diag1|pos)<<1,
            (diag2|pos)>>1,
            n,
        )
    }

    return count
}

func main() {
    var n int

    if _, err := fmt.Scan(&n); err != nil {
        return
    }

    total := solve(0, 0, 0, 0, n)

    fmt.Println(total)
}