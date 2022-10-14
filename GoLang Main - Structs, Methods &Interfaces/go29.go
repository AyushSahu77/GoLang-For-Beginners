package main

import "fmt"

type shape interface { // Interface Syntax : type <name> interface
	area() float64
	perimeter() float64
}

type rect struct {
	length  float64
	breadth float64
}

type square struct {
	side float64
}

func (re rect) area() float64 { // we can say that rect has implemented shape interface
	return (re.length * re.breadth)
}

func (re rect) perimeter() float64 {
	return 2 * (re.length + re.breadth)
}

func (sq square) area() float64 { // we can say that square has implemented shape interface
	return (sq.side * sq.side)
}

func (sq square) perimeter() float64 {
	return (4 * sq.side)
}

func displayData(s shape) { // one single method to print all func that are inside interface shape
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	re := rect{
		length:  4,
		breadth: 7,
	}

	sq := square{
		side: 5,
	}

	displayData(re)
	displayData(sq)
}
