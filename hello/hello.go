package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a, b := 1, 100
	r := rand.New(rand.NewSource(99))
	secret := r.Intn(100)
	fmt.Printf("%d is guessed\n", secret)
	for {
		fmt.Printf("Input number from %d to %d\n", a, b)
		var guess int
		fmt.Scan(&guess)
		if guess == secret {
			fmt.Println("You guessed!")
			return
		}
		if guess < secret {
			fmt.Println("Less")
		} else {
			fmt.Println("More")
		}
	}
}
