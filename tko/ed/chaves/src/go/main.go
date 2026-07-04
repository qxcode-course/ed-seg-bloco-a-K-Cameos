package main

import "fmt"

func main() {
	q := NewQueue[string]()

	for i := 0; i < 16; i++ {
		team := string(rune('A' + i))
		q.Enqueue(team)
	}

	for !q.IsEmpty() {
		if q.items.Len() == 1 {
			break
		}
		
		time1 := q.Dequeue()
		time2 := q.Dequeue()

		var gols1, gols2 int

		if _, err := fmt.Scan(&gols1, &gols2); err != nil {
			break
		}

		if gols1 > gols2 {
			q.Enqueue(time1)
		} else {
			q.Enqueue(time2)
		}
	}

	campeao := q.Dequeue()
	fmt.Println(campeao)
}