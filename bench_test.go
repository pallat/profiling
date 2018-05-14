package main

import (
	"fmt"
	"testing"
)

func TestNonsense(t *testing.T) {
	nonsense(10)
}

// func BenchmarkNonsense(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		nonsense(1)
// 	}
// }

func Fib(n int) {
	fmt.Println(n)
}

func BenchmarkFibWrong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(n)
	}
}
