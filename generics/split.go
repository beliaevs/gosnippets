package main

import (
	"fmt"
)

// very weird way to allow comparison in generic
type Ordered interface {
	~int | ~float32 | ~float64 | ~string
}

func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func findMax[T Ordered](s []T) (res T, err error) {
	if len(s) == 0 {
		err = fmt.Errorf("findMax: empty slice")
		return
	}
	res = s[0]
	for _, e := range s[1:] {
		if res < e {
			res = e
		}
	}
	return
}

func main() {
	txt := [...]string{"hello", "goodbye", "hi", "bye", "sasha"}
	fmt.Println(txt)
	half1, half2 := splitAnySlice(txt[:])
	fmt.Println(half1, half2)

	m, err := findMax(txt[:])
	if err == nil {
		fmt.Println(m)
	} else {
		fmt.Println(err)
	}
}
