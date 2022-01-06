package main

import (
	"fmt"
	"time"
)

func main() {
	const size = 20000
	var twoDimensions = [size][size]uint32{}

	var start = time.Now()

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			twoDimensions[i][j]++
		}
	}
	var elapsedFast = time.Since(start)

	fmt.Println("FAST: ")
	fmt.Println(elapsedFast)

	twoDimensions = [size][size]uint32{}
	start = time.Now()
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			twoDimensions[j][i]++
		}
	}

	var elapsedSlow = time.Since(start)

	fmt.Println("SLOW: ")
	fmt.Println(elapsedSlow)

	fmt.Println("Speedup: ")
	var speedup = float64(elapsedFast.Nanoseconds()) / float64(elapsedSlow.Nanoseconds())
	fmt.Printf("%.2fx\n", 1.0/(speedup))
}
