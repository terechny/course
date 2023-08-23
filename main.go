package main

import (
	"course/buz"
	"fmt"
	"math"
)

func main() {

	for i := 0; i < 2; i++ {

		fmt.Println(i, buz.Buzz(i))
	}

	fmt.Println(summ(10, 20))
	fmt.Println(summs(45))

	number := 12.3456789
	fmt.Println(roundFloat(number, 2))
}

func summ(a, b int) int {
	return a + b
}
func summs(n int) int {
	s := 0
	for i := 0; i <= n; i++ {

		s += i
	}
	return s
}

func roundFloat(val float64, p uint) float64 {
	r := math.Pow(10, float64(p))
	return math.Round(val*r) / r
}
