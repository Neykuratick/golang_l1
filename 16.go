package main

import (
	"fmt"
	"math/rand"
	"time"
)

func merge(slice1 []int, pivot int, slice2 []int) []int {
	var masterSlice []int

	for _, num := range slice1 {
		masterSlice = append(masterSlice, num)
	}

	masterSlice = append(masterSlice, pivot)

	for _, num := range slice2 {
		masterSlice = append(masterSlice, num)
	}

	return masterSlice
}

func quicksort(numbers []int) []int {
	if len(numbers) < 1 {
		return numbers
	}

	var greater []int
	var smaller []int
	pivot := numbers[len(numbers)-1:][0]
	numbers = numbers[:len(numbers)-1]

	for _, number := range numbers {
		if number > pivot {
			greater = append(greater, number)
		} else {
			smaller = append(smaller, number)
		}
	}

	return merge(quicksort(smaller), pivot, quicksort(greater))
}

func main() {
	var numbers []int
	rand.Seed(time.Now().Unix())

	for i := 0; i < 15; i++ {
		numbers = append(numbers, rand.Intn(10000))
	}

	now := time.Now()
	numbers = quicksort(numbers)
	fmt.Println(numbers)
	fmt.Println(time.Now().Sub(now).Seconds())
}
