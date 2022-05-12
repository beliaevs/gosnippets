package main

import (
	"fmt"
)

func main() {
	s := 0.0
	N := 100000000
	for i := 1; i < N; i += 1 {
		s += 1.0 / float64(N-i)
	}
	fmt.Println("s = ", s)
}
