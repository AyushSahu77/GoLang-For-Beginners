package main

import "fmt"

type circle struct {
	radius float64
	area   float64
}

func (c *circle) circleArea() { // struct method Syntax : func (struct pointer) name (parameter)
	const pi = 3.14
	var area float64 = pi * c.radius * c.radius
	c.area = area
}

func main() {
	c := circle{
		radius: 7,
	}
	c.circleArea()
	fmt.Println(c.area)
}
