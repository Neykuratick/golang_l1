package main

import "fmt"

func setBitTo0(n int64, pos uint) int64 {
	return n &^ (1 << pos)
}

func setBitTo1(n int64, pos uint) int64 {
	n |= 1 << pos
	return n
}

func main() {
	n := int64(10)
	fmt.Printf("Original: %b\n\n", n)
	n = setBitTo1(n, 2)
	fmt.Printf("Set bit: %b\n", n)
	n = setBitTo0(n, 2)
	fmt.Printf("Clear bit: %b\n", n)
}
