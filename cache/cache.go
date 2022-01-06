package main

import (
	"fmt"
	"time"
)

func main() {
	const size = 20000

	array := createSlice(size)
	var elapsedFast = runGoodCaching(array)
	fmt.Println("FAST: ")
	fmt.Println(elapsedFast)

	array = createSlice(size)
	var elapsedSlow = runBadCaching(array)
	fmt.Println("SLOW: ")
	fmt.Println(elapsedSlow)

	fmt.Println("Speedup: ")
	var speedup = float64(elapsedFast.Nanoseconds()) / float64(elapsedSlow.Nanoseconds())
	fmt.Printf("%.2fx\n", 1.0/(speedup))
}

func createSlice(size int) [][]uint32 {
	twoDimensions := make([][]uint32, size)
	for i := range twoDimensions {
		twoDimensions[i] = make([]uint32, size)
	}

	return twoDimensions
}

func runGoodCaching(array [][]uint32) time.Duration {
	var size = len(array)
	var start = time.Now()

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			array[i][j]++
		}
	}

	return time.Since(start)
}

func runBadCaching(array [][]uint32) time.Duration {
	var size = len(array)
	var start = time.Now()

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			array[j][i]++
		}
	}

	return time.Since(start)
}
