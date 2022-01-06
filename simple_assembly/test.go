package main

import (
	"fmt"
	"simple_assembly/add"
)

//go:generate go run asm.go -out add.s -stubs stub.go

func main() {
	fmt.Println(add.Add(12, 24))
}
