package main

import (
	"course/buz"
	"testing"
)

/*
	func BenchmarkSumms(b *testing.B) {
		for i := 0; i < b.N; i++ {
			summs(10)
		}
	}

	func BenchmarkSumm(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			summ(10, 20)
		}
	}
*/

func TestSm(t *testing.T) {

	testCases := []struct {
		name string
		a    int
		b    int
		out  int
	}{
		{
			name: "1+1",
			a:    1,
			b:    1,
			out:  2,
		},
		{
			name: "562 + 960",
			a:    562,
			b:    960,
			out:  15220,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			got := summ(tc.a, tc.b)

			if got != tc.out {
				t.Errorf("expected %d, got %d", tc.out, got)
			}
		})
	}
}

func TestFizzBuff(t *testing.T) {

	testCases := []struct {
		name string
		in   int
		out  string
	}{
		{
			name: "3",
			in:   3,
			out:  "fiz",
		},
		{
			name: "5",
			in:   5,
			out:  "buzz",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			got := buz.Buzz(tc.in)
			if got != tc.out {
				t.Errorf("fizzBuzz(%d) = %s; expected %s", tc.in, got, tc.out)
			}
		})
	}
}

func BenchmarkFibonacciRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buz.FibonacciRecursion(20)
	}
}

func BenchmarkFibonacciBinet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buz.FibonacciBinet(20)
	}
}
