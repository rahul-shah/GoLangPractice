package main

import "fmt"

type MyPoint struct {
	X int
	Y int
}

type Square struct {
	Centre MyPoint
	Length int
}

func (s *Square) Move(dx int, dy int) {
	s.Centre.X += dx
	s.Centre.Y += dy
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("Value of length can't be less than 0")
	}
	p := MyPoint{x, y}
	s := Square{p, length}
	return &s, nil
}

func main() {
	s, error := NewSquare(1, 2, 5)
	if error == nil {
		fmt.Printf("%+v\n", s.Area())
	} else {
		fmt.Printf("An Error has occured\n")
	}

}
