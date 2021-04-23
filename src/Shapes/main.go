package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLenght float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	sq := square{5.0}
	tr := triangle{5.0, 3.0}

	printArea(sq)
	printArea(tr)
}

func printArea(sh shape) {

	fmt.Println(sh.getArea())
}

func (sq square) getArea() float64 {
	return sq.sideLenght * sq.sideLenght
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}
