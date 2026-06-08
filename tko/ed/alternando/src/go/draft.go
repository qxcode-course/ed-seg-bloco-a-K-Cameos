package main
import "fmt"
func main() {
    var n, e, f int
	if _, err := fmt.Scan(&n, &e, &f); err != nil {
		return
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = (i + 1) * f
		f = -f
	}

	swordIdx := e - 1

	for len(arr) > 1 {
		printState(arr, swordIdx)

		v := arr[swordIdx]
		dir := 1
		if v < 0 {
			dir = -1
		}

		killIdx := (swordIdx + dir + len(arr)) % len(arr)

		arr = append(arr[:killIdx], arr[killIdx+1:]...)

		if killIdx < swordIdx {
			swordIdx--
		}
		swordIdx = (swordIdx + dir + len(arr)) % len(arr)
	}
printState(arr, swordIdx)

}

func printState(arr []int, swordIdx int) {
	fmt.Print("[ ")
	for i, v := range arr {
		if i == swordIdx {
			if v > 0 {
				fmt.Printf("%d> ", v)
			}else {
				fmt.Printf("<%d ", v)
			}
		} else {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println("]")
}