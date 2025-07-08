package main

import (
	"fmt"
	"goexcercises/norepeats"
	"time"
)

func main() {

	sentence := "candy canes do taste yummy"
	start := time.Now()

	lastUniqueChar := norepeats.NonRepeat(sentence)

	elapsed := time.Since(start)

	fmt.Println("Lasr non-repeat:", string(lastUniqueChar))

	fmt.Printf("Function took: %v\n", elapsed)
}
