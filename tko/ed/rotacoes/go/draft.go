package main
import "fmt"
func main() {
    var t, r int
	
	if _, err := fmt.Scan(&t, &r); err != nil {
		return
	}
	if t == 0 {
		fmt.Println("[]")
		return
	}

	arr := make([]int, t)
	for i := 0; i<t; i++ {
		fmt.Scan(&arr[i])
	}

	r = r % t
	split := (t-r) % t

	rot := append(arr[split:], arr[:split]...)

	fmt.Print("[ ")
	for _, val := range rot {
		fmt.Printf("%d ", val)
	}
	fmt.Println("]")
}