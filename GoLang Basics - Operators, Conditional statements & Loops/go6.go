package main

import "fmt"

func main() {
	var a int = 4
	var b int = 7
	var c int = 8
	var s1 string = "Bhilai"
	var s2 string = "Durg"
	var s3 string = "Bhilai"
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(c >= b)
	fmt.Println("\n\n")
	fmt.Println(s1 == s2)
	fmt.Println(s1 == s3)
	fmt.Println(s1 != s3)
}
