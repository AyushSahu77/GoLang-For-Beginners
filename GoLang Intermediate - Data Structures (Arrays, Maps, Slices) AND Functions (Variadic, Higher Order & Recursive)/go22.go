package main

import (
	"fmt"
)

func printSem(sem int) {
	fmt.Println(sem, "th")
}

func printName(name string) {
	fmt.Println(name)
}

func printrollno(roll int) {
	fmt.Println(roll)
}
func printsection(sec string) {
	fmt.Println(sec)
}

func main() {
	printSem(5)
	defer printName("Ayush") //defer delays the function output untill all function after that are not executed
	defer printrollno(54)    // last in first out functionality is used in defer
	printsection("B")
}
