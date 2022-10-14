package main

import (
	"fmt"
)

func main() {
	var arr1 [3]int = [3]int{10, 20, 30}
	var arr2 [5]int
	arr3 := [4]int{100, 200, 300, 400}       //short hand declaration
	arr4 := [...]string{"aaa", "bbb", "ccc"} //elipse declaration
	arr5 := [...]int{12, 34, 56, 78, 90}
	var arr2d1 [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < 5; i++ {
		arr2[i] = i //assigning
	}
	fmt.Println(arr2) // printing
	fmt.Println(arr1) // printing all like array
	for l := 0; l < 3; l++ {
		fmt.Println(arr3[l]) //printing each element
	}
	for m := 0; m < len(arr4); m++ {
		fmt.Println(arr4[m]) //string array
	}
	fmt.Println("length of arr4 = ", len(arr4))

	for index, element := range arr5 { //range is a function through which we get index and element's values of the array
		fmt.Println(index, "=>", element)
	}
	fmt.Println(arr2d1) // simple 2d array output
	for index, value := range arr2d1 {
		fmt.Println(index, "=>", value) // using range func
	}
}
