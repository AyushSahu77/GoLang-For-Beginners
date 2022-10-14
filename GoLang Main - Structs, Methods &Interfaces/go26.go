package main

import (
	"fmt"
)

type rect struct {
	l    float64
	b    float64
	area float64
}

func rectArea(r *rect) {
	var area float64 = r.l * r.b
	(*r).area = area
}

func main() {
	r := rect{
		l: 10,
		b: 20,
	}
	fmt.Printf("%+v \n", r)
	rectArea(&r) // we have to pass values by reference in order to change in the struct
	fmt.Println(r.area)

}
