package main

import (
	"fmt"
)

func main() {
	var a string = "Ayush"
	if a == "Ayush" {
		fmt.Println("Yes")
	} else if a == "ayush" {
		fmt.Println("partially")
	} else {
		fmt.Println("No")
	}
}
