package main

import (
	"fmt"
	"math/big"
)

func divide(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Div(a, b)
	return result
}

func multiply(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Mul(a, b)
	return result
}

func add(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Add(a, b)
	return result
}

func subtract(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Sub(a, b)
	return result
}

func main() {
	a := new(big.Int)
	a.SetString("999999999999999999999999999999999999999999999999999999", 10)

	b := new(big.Int)
	b.SetString("333333333333333333333333333333333333333333333333333333", 10)

	fmt.Printf("a: %v\nb: %v\n\n", a, b)
	fmt.Println("Divide:", divide(a, b))
	fmt.Println("Multiply:", multiply(a, b))
	fmt.Println("Add:", add(a, b))
	fmt.Println("Subtract:", subtract(a, b))
}
