package main

import (
	"fmt"
)

func main() {
	limit := 4000000
	sum := 0
	a := 1
	b := 1
	for i := 1; b < limit; i++ {
		temp := a + b
		a = b
		b = temp
		if b%2 == 0 {
			sum = sum + b
		}
	}
	fmt.Println(sum)
}
