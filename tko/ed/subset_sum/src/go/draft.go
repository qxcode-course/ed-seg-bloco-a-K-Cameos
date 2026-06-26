package main
import "fmt"

func subsetSumDP(arr []int, target int) bool {
    dp := make([]bool, target+1)
    dp[0] = true

    for  _, num := range arr {
        for j := target; j >= num; j-- {
            if dp[j-num] {
                dp[j] = true
            }
        }
    }

    return dp[target]

}

func main() {
    var n, k int

    if _, err := fmt.Scan(&n, &k); err != nil {
        return
    }
    arr := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    fmt.Println(subsetSumDP(arr, k))
}