package main

import (
	"fmt"
	"math/rand"
	"net"
)

type figure struct {
	XPos   int
	YPos   int
	Width  int
	Height int
	Color  string
}

var brick []figure
var ball figure
var paddle figure
var connection net.Conn

func render() {
	fmt.Fprintf(connection, "clear\n")

	for i := range brick {
		renderFigure(brick[i])
	}

	renderFigure(ball)
	renderFigure(paddle)

	fmt.Fprintf(connection, "close\n")
}

func renderFigure(f figure) {
	var text string
	for y := f.YPos; y < f.YPos+f.Height; y++ {
		if y%2 == 0 {
			text = fmt.Sprintf("MEM %x", (f.XPos + (y/2)*80))
		} else {
			text = fmt.Sprintf("MEM %x", ((f.XPos + (y/2)*80) + 8000))
		}

		for x := f.XPos; x < f.XPos+f.Width; x++ {
			text = fmt.Sprintf("%s %s", text, f.Color)
		}

		fmt.Fprintf(connection, "%v\n", text)
	}
}

func createFigures() {
	col := [3]string{"55", "AA", "FF"}

	for y := 0; y < 6; y++ {
		for x := 0; x < 11; x++ {
			posX := x*5 + (x+1)*2
			posY := y*4 + (y+1)*8
			rand := rand.Intn(len(col))
			brick = append(brick, figure{XPos: posX, YPos: posY, Width: 5, Height: 5, Color: col[rand]})
		}
	}

	fmt.Printf("created %d\n", brick)

	ball = figure{XPos: 42, YPos: 100, Width: 1, Height: 8, Color: "FF"}

	paddle = figure{XPos: 10, YPos: 185, Width: 10, Height: 5, Color: "FF"}
}

func main() {

	createFigures()
	connection, _ = net.Dial("tcp", "localhost:1337")

	fmt.Println("start rendering...")

	fmt.Fprintf(connection, "reg 0x03D8 2\n")
	render()

	//fmt.Fprintf(connection, "close\n")

}
