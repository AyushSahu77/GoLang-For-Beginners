package main

import (
	"fmt"
)

const pi float64 = 3.14

func circleArea(r float64) float64 {
	return pi * r * r
}

func circleCircumference(r float64) float64 {
	return 2 * pi * r
}

func circleDia(r float64) float64 {
	return 2 * r
}

func higherOrderFunc1(r float64, calcFunc func(r float64) float64) {
	result := calcFunc(r)
	fmt.Println("result : ", result)
} // higher order func that takes a func as an arguement

func higherOrderFunc2(option int) func(r float64) float64 {
	optionChoosed := map[int]func(r float64) float64{
		1: circleArea,
		2: circleCircumference,
		3: circleDia,
	}
	return optionChoosed[option]
} // higher order func that returns a func as an arguement

func main() {
	/* var r float64
	var option int
	fmt.Println("Enter radius of the circle")
	fmt.Scanf("%f", &r)
	fmt.Println("What to do \n 1 : Area \n 2 : Circumference \n 3 : Dia")
	fmt.Scanf("%d", &option)
	if option == 1 {
		fmt.Println("Area : ", circleArea(r))
	} else if option == 2 {
		fmt.Println("circumference : ", circleCircumference(r))
	} else if option == 3 {
		fmt.Println("Dia : ", circleDia(r))
	} else {
		fmt.Println("wrong option")
	} */
	var r float64
	var option int
	fmt.Println("Enter radius of the circle")
	fmt.Scanf("%f", &r)
	fmt.Println("What to do \n 1 : Area \n 2 : Circumference \n 3 : Dia")
	fmt.Scanf("%d", &option)
	higherOrderFunc1(r, higherOrderFunc2(option))
}
