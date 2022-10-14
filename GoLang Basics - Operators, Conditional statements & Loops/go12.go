package main

import (
	"fmt"
)

func main() {
	var s string = "Ayush"
	switch s {
	case "Ayush":
		fmt.Println("Ayush")
		fallthrough //forces us to execute next switch case
	case "ayush":
		fmt.Println("ayush")
	default:
		fmt.Println("Default")
	}
	var a, b int = 3, 4
	switch {
	case a+b == 7:
		fmt.Println("=7")
		fallthrough
	case a+b <= 7: // this block won't execute, GoLang follows implicit break as soon as a case is true, we'll have to use fallthrough keyword to continue, this is very diff from c++ and java
		fmt.Println("<=7")
	case a+b > 7:
		fmt.Println(">7")
		fallthrough //  this doesn't work, as the case was false the control never came inside the case block
	case b-a == 1:
		fmt.Println("b-a = 1")
	default:
		fmt.Println(".")
	}
}
