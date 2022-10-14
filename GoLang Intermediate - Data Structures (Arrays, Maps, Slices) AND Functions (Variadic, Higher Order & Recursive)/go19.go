package main

import "fmt"

func variadicFunc1(numbers ...int) (sum int) { //single arguemnt
	sum = 0
	for _, value := range numbers {
		sum += value
	}
	fmt.Println(sum)
	return
}

func variadicFunc2(month string, day ...string) (mon string, da string) {
	fmt.Println(month) // month is mandatory arguement, day can be from 0 to any number
	for _, day := range day {
		fmt.Printf("%s, ", day)
	}
	return
}

func variadicFunc3() (int, int) {
	return 4, 9
}

func main() {
	variadicFunc1(1, 2, 3, 4)
	fmt.Print("\n")
	variadicFunc2("september", "1", "2", "3")
	fmt.Println()
	fmt.Println()
	p, _ := variadicFunc3() // a _ can be used if that value is not needed
	fmt.Println(p)          // only one value printed
}
