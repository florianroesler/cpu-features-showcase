package main

import (
	"fmt"
	"time"
)

func main() {
	var iterations = 1500000000

	var start = time.Now()
	var sum = 0
	for i := 0; i < iterations; i++ {
		if i&128 == 0 {
			sum++
		}
	}
	var elapsedFast = time.Since(start)

	fmt.Println("128T 128F: ")
	fmt.Println(sum)
	fmt.Println(elapsedFast)

	start = time.Now()
	sum = 0
	for i := 0; i < iterations; i++ {
		if i&2 == 0 {
			sum++
		}
	}
	var elapsedSlow = time.Since(start)

	fmt.Println("1T 1F: ")
	fmt.Println(sum)
	fmt.Println(elapsedSlow)

	fmt.Println("Speedup: ")
	var speedup = float64(elapsedFast.Nanoseconds()) / float64(elapsedSlow.Nanoseconds())
	fmt.Printf("%.2fx\n", 1.0/(speedup))
}
