package main

import (
	"fmt"
	"time"
)

func main() {
	var iterations = 1_000_000_000

	var fastBenchmarkTime = runStableConditional(iterations)
	fmt.Println(fastBenchmarkTime)

	var slowBenchmarkTime = runOscillatingConditional(iterations)
	fmt.Println(slowBenchmarkTime)

	fmt.Println("Speedup: ")
	var speedup = float64(fastBenchmarkTime.Nanoseconds()) / float64(slowBenchmarkTime.Nanoseconds())
	fmt.Printf("%.2fx\n", 1.0/(speedup))
}

func runStableConditional(iterations int) time.Duration {
	var start = time.Now()
	var sum = 0
	for i := 0; i < iterations; i++ {
		if i&128 == 0 {
			sum++
		}
	}

	fmt.Println("128T 128F: ")
	fmt.Println(sum)

	return time.Since(start)
}

func runOscillatingConditional(iterations int) time.Duration {
	var start = time.Now()
	var sum = 0
	for i := 0; i < iterations; i++ {
		if i&1 == 0 {
			sum++
		}
	}

	fmt.Println("1T 1F: ")
	fmt.Println(sum)

	return time.Since(start)
}
