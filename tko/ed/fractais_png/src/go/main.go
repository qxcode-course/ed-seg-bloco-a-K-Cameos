package main

import (
	"fmt"
)

func drawGelo(pen *Pen, length float64, depth int, isRoot bool) {
	if depth == 0 {
		return
	}

	reducao := 0.35

	if isRoot {
		for i := 0; i < 5; i++ {
			pen.Down()
			pen.Walk(length)

			drawGelo(pen, length*reducao, depth-1, false)

			pen.Up()
			pen.Walk(-length)
			pen.Left(72)
		}
	} else {
		pen.Left(108)

		for i := 0; i < 4; i++ {
			pen.Down()
			pen.Walk(length)

			drawGelo(pen, length*reducao, depth-1, false)

			pen.Up()
			pen.Walk(-length)
			pen.Right(72)
		}

		pen.Left(108)
	}
}

func main() {
	pen := NewPen(800, 800)

	pen.SetRGB(0, 0, 0)
	pen.dc.Clear()

	pen.SetRGB(255, 255, 255)
	pen.SetLineWidth(1)

	pen.Up()
	pen.SetPosition(400, 400)
	pen.SetHeading(270)
	pen.Down()

	raioInicial := 230.0
	profundidade := 5

	drawGelo(pen, raioInicial, profundidade, true)

	pen.SavePNG("gelo.png")
	fmt.Println("png criado: gelo.png")
}
