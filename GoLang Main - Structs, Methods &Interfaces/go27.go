package main

import (
	"fmt"
)

type s1 struct {
	a int
}

type s2 struct {
	a int
}

func main() {
	o1 := s1{
		a: 4,
	}

	o2 := s1{
		a: 4,
	}

	o3 := s1{
		a: 7,
	}
	/*xs1 := s1 {
		a:3,
	}
	xs2 := s2 {
		a:3,
	}*/
	// xs1 == xs2 {   we can't compare like this as both are of diff types, xs1 in of type s1 and xs2 is of type s2
	if o1 == o2 {
		fmt.Println("Equal")
	}
	if o1 == o3 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
}
