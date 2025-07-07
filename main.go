package main


import (
	"go_excercises/fireworks"
	"fmt"
	"time"
)

func main() {
	fireworkData := []fireworks.Firework{
		{Height: 10, Size: 6, Velocity: 4},
		{Height: 13, Size: 3, Velocity: 2},
		{Height: 17, Size: 6, Velocity: 3},
		{Height: 21, Size: 8, Velocity: 4},
		{Height: 19, Size: 5, Velocity: 3},
		{Height: 18, Size: 4, Velocity: 4},
	}

	start := time.Now()

	startIndex := fireworks.GetFinaleStart(fireworkData)

	elapsed := time.Since(start)

	fmt.Println("The finale starts at index:", startIndex)
	fmt.Printf("Function took: %v\n", elapsed)
}