package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partition(numbers []int, start int, end int) int {
	pivot := numbers[end]
	j := start - 1 // можно начинать и с 0, но тогда надо будет делать j++ после свопа

	for i := start; i < end; i++ {
		if numbers[i] < pivot {
			j++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}

	j++
	numbers[end], numbers[j] = numbers[j], numbers[end]

	return j
}

func Quicksort16(numbers []int, start int, end int) []int {
	// быстрая сортировка, алгоритм которой я посмотрел на ютубе
	// она быстрее, потому что работает с одним только слайсом

	if start > end {
		return numbers
	}
	pivot := partition(numbers, start, end)
	numbers = Quicksort16(numbers, start, pivot-1)
	numbers = Quicksort16(numbers, pivot+1, end)
	return numbers
}

func main() {
	var numbers []int
	rand.Seed(time.Now().Unix())

	for i := 0; i < 15; i++ {
		numbers = append(numbers, rand.Intn(10000))
	}

	now := time.Now()
	numbers = Quicksort16(numbers, 0, len(numbers)-1)
	fmt.Println(numbers)
	fmt.Println(time.Now().Sub(now).Milliseconds()) // тут ещё замеры есть
}
