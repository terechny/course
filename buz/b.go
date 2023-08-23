package buz

import (
	"math"
	"strconv"
)

func Buzz(n int) string {

	if n%3 == 0 && n%5 == 0 {
		return "fizz buzz"
	}

	if n%3 == 0 {
		return "fiz"
	}

	if n%5 == 0 {
		return "buzz"
	}

	return strconv.Itoa(n)
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func FibonacciBinet(n int) int {
	sqrt5 := math.Sqrt(5)
	p := (1 + sqrt5) / 2
	q := 1 / p
	return int((math.Pow(p, float64(n)) - math.Pow(q, float64(n))) / sqrt5)
}
