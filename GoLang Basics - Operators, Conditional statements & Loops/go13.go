package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println(i * i * i)
	}
	j := 0
	for j <= 4 {
		fmt.Println(j * j)
		j += 1
	}
	for k := 0; k < 10; k++ {
		if k == 7 {
			fmt.Println("Broke")
			break
		}
		fmt.Println(k)
	}
	for l := 0; l < 10; l++ {
		if l == 7 {
			fmt.Println("Continued skipping this iteration ", l)
			continue
		}
		fmt.Println(l)
	}
}
