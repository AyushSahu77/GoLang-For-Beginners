package main

import (
	"fmt"
)

func main() {
	arr1 := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice1 := arr1[0:8]
	subslice1 := slice1[3:5]
	slice2 := make([]int, 4, 7) // Syntax : make(datatype,len,cap)

	fmt.Println(slice1)
	fmt.Println(subslice1)
	fmt.Println(slice2)
	fmt.Print("\n")

	arr2 := [5]int{1, 2, 3, 4, 5}
	slice3 := arr2[:3]
	fmt.Println(arr2)
	fmt.Println(slice3)
	slice3[1] = 70
	fmt.Println("Modified slice value affects array as well") //as it is a reference to the array
	fmt.Println(arr2)
	fmt.Println(slice3)
	fmt.Print("\n")

	arr3 := [5]int{1, 2, 3, 4, 5}
	slice4 := arr3[1:4]
	fmt.Println(slice4)
	fmt.Println(len(slice4)) //length of the slice or array
	fmt.Println(cap(slice4)) // capacity of the slice will 4 as starting from 1 it goes till 2 3 4 means cap = 4, won't count 0 as slice starts from 1
	// for array the capacity = length
	slice5 := append(slice4, 70) //appending slice with another slice value + extra values
	fmt.Println(slice5)
	fmt.Println(len(slice5)) // length increases by 1 as 1 element is appended
	fmt.Println(cap(slice5)) // capacity remains same as there was space to fit the extra appended element
	slice6 := append(slice4, 80, 90, 100)
	fmt.Println(slice6)
	fmt.Println(len(slice6)) // length increases by 3 as 3 element is appended
	fmt.Println(cap(slice6)) // capacity DOUBLES as there was not enough space to fit the extra appended elements
	//	LENGTH INCREASES BY NO. OF EXTRA ELEMENTS AND CAPACITY DOUBLES IF NOT ENOUGH
	fmt.Print("\n")
	arr4 := []int{1, 2, 3, 4, 5}
	slice7 := arr4[1:3]
	arr5 := []int{6, 7, 8, 9, 0}
	slice8 := arr5[1:4]
	slice9 := append(slice7, slice8...)
	fmt.Println(slice9)
	fmt.Println(len(slice9))
	fmt.Println(cap(slice9))
	fmt.Print("\n")
	src_slice := []int{10, 20, 30, 40, 50}
	dst_slice := make([]int, 3)       //when only specified 1 arguement, it is length, capacity is not must, and cap is assigned same as len
	num := copy(dst_slice, src_slice) // Syntax : copy(destination slice, source slice)  ||||||  num stores the number of elemnts copied
	fmt.Println(dst_slice)
	fmt.Println("Num of elmt", num)

	//Looping of slice is same as array with range func
	// if index is not needed in range func we can write as

	/*
		for _,value := range arrayname {
			fmt.Println(value)
			//an underscore is used inplace of index variable
		}
	*/
}
