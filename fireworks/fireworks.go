package fireworks

import "math"

type Firework struct {
	Height   int
	Size     int
	Velocity int
}

func GetFinaleStart(fireworks []Firework) int {
	startIndex := -1
	maxLength := 0

	for i := 0; i < len(fireworks); i++ {
		minHeight := math.MaxInt
		maxHeight := math.MinInt
		minVelocity := math.MaxInt
		sumSize := 0

		for j := i; j < len(fireworks); j++ {
			minHeight = min(minHeight, fireworks[j].Height)
			maxHeight = max(maxHeight, fireworks[j].Height)
			minVelocity = min(minVelocity, fireworks[j].Velocity)

			sumSize += fireworks[j].Size
			currentLength := (j - i) + 1
			avgSize := sumSize / currentLength

			if avgSize >= 5 && minVelocity >= 3 && maxHeight-minHeight <= 10 {
				if currentLength > maxLength {
					startIndex = i
					maxLength = currentLength
				}
			}
		}
	}
	return startIndex
}
