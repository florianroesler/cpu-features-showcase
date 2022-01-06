package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bjwbell/gensimd/simd"
)

func main() {
	rand.Seed(1000)

	var iterations = 500000000

	var start = time.Now()
	var a = [4]int32{rand.Int31(), rand.Int31(), rand.Int31(), rand.Int31()}
	var b = [4]int32{rand.Int31(), rand.Int31(), rand.Int31(), rand.Int31()}
	var c = [4]int32{0, 0, 0, 0}

	for i := 0; i < iterations; i++ {
		addSimple(a, b, c)
	}
	var elapsedNoSSE = time.Since(start)

	fmt.Println("NO SSE: ")
	fmt.Println(elapsedNoSSE)

	start = time.Now()

	var aSimd simd.I32x4
	aSimd[0] = rand.Int31()
	aSimd[1] = rand.Int31()
	aSimd[2] = rand.Int31()
	aSimd[3] = rand.Int31()

	var bSimd simd.I32x4
	bSimd[0] = rand.Int31()
	bSimd[1] = rand.Int31()
	bSimd[2] = rand.Int31()
	bSimd[3] = rand.Int31()

	for i := 0; i < iterations; i++ {
		add(aSimd, bSimd)
	}
	var elapsedSSE = time.Since(start)

	fmt.Println("SSE: ")
	fmt.Println(elapsedSSE)

	fmt.Println("Speedup: ")
	var speedup = float64(elapsedSSE.Nanoseconds()) / float64(elapsedNoSSE.Nanoseconds())
	fmt.Printf("%.2fx\n", 1.0/(speedup))
}

func addSimple(a [4]int32, b [4]int32, c [4]int32) [4]int32 {
	c[0] = a[0] + b[0]
	c[1] = a[1] + b[1]
	c[2] = a[2] + b[2]
	c[3] = a[3] + b[3]
	return c
}

func add(x, y simd.I32x4)

//
//
//
//
//
// func addi32x4(x, y simd.I32x4) simd.I32x4 {
// 	return simd.AddI32x4(x, y)
// }
