package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	p := Point{1, 2}
	p.Move(3, 4)
	fmt.Printf("%+x\n", p)
}
