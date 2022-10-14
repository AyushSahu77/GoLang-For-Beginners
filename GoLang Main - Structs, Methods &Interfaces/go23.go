package main

import (
	"fmt"
)

func main() {
	var ptr1 *int
	i := 10
	j := 20
	s := "Hello"
	f := 3.4
	fmt.Printf("%T %v \n", &i, &i)       // prints type and value of pointer
	fmt.Printf("%T %v \n", *(&i), *(&i)) //prints type and value of dereferenced pointer
	fmt.Println(ptr1)                    // o/p = <nil> as zero value of pointer is nil
	var ptr2 *int = &i                   // storing address manually
	fmt.Println(ptr2)
	var ptr3 = &s
	fmt.Println(ptr3) // datatype will be internally determined by the compiler i.e. string
	ptr4 := &f
	fmt.Println(ptr4) // shorthand declaration of pointers
	fmt.Println()
	fmt.Println("original value of j", j)
	ptr5 := &j
	fmt.Println(ptr5)
	*ptr5 = 30 // dereferencing pointer and changing value
	fmt.Println("new changed value of j", j)
	ptr6 := &j // address will be same as only the value has changed the position hasn't
	fmt.Println(ptr6)
}
