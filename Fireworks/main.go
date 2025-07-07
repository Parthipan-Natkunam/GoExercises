package main

import (
	"fmt"
	"math"
	"time"
)

type Firework struct {
	height   int
	size     int
	velocity int
}

func getFinaleStart(fireworks []Firework) int {
	startIndex := -1
	maxLength := 0

	for i := 0; i < len(fireworks); i++ {
		minHeight := math.MaxInt
		maxHeight := math.MinInt
		minVelocity := math.MaxInt
		sumSize := 0

		for j := i; j < len(fireworks); j++ {
			minHeight = min(minHeight, fireworks[j].height)
			maxHeight = max(maxHeight, fireworks[j].height)
			minVelocity = min(minVelocity, fireworks[j].velocity)

			sumSize += fireworks[j].size
			currentLength := (j - i) + 1
			avgSize := sumSize / currentLength

			if avgSize >= 5 && minVelocity >= 3 && maxHeight-minHeight <=10 {
				if currentLength > maxLength {
					startIndex = i
					maxLength = currentLength
				}
			}
		}
	}
	return startIndex
}

func main() {
	fireworks := []Firework{
		{height: 10, size: 6, velocity: 4},
		{height: 13, size: 3, velocity: 2},
		{height: 17, size: 6, velocity: 3},
		{height: 21, size: 8, velocity: 4},
		{height: 19, size: 5, velocity: 3},
		{height: 18, size: 4, velocity: 4},
	}

	start := time.Now()

	startIndex := getFinaleStart(fireworks)

	elapsed := time.Since(start)

	fmt.Println("The finale starts at index:", startIndex)
	fmt.Printf("Function took: %v\n", elapsed)
}
