package main

import (
	"fmt"
)

type student struct {
	name    string
	sem     int
	rollno  int
	section string
	marks   []int
}

func main() {
	//var s student
	//fmt.Printf("%+v", s)
	s := student{
		name:    "Ayush",
		sem:     5,
		rollno:  54,
		section: "B",
		marks:   []int{10, 20, 30, 40},
	}
	fmt.Printf("%+v \n", s)
	fmt.Println(s.name) // we can print sepcific fields of struct and we can also assign them this way
	s.name = "Sahu"
	fmt.Println(s.name) //changed value
}
