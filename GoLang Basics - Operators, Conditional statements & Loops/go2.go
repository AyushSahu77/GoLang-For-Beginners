package main

import "fmt"

func main() {
	var a string
	var b int
	//var name string
	// var rollno int
	fmt.Println("Enter your a and b")
	count, err := fmt.Scanf("%s %d", &a, &b)
	fmt.Println("count", count)
	fmt.Println("err", err)
	fmt.Println("a", a)
	fmt.Println("b", b)
	// fmt.Println("Enter your name and roll no")
	// fmt.Scanf("%s %d", &name, &rollno)
	// fmt.Println(surname, rollno)
}
