package main

import "fmt"

func binarySearch(numbers []int, number int) bool {
	// стринги можно так же искать, если их в руны конвентировать

	if len(numbers) < 2 {
		return numbers[0] == number
	}

	middle := len(numbers) / 2
	last := len(numbers) - 1

	numbers = Quicksort17(numbers, 0, last)

	if numbers[middle] > number {
		return binarySearch(numbers[:middle], number)
	} else if numbers[middle] < number {
		return binarySearch(numbers[middle+1:], number)
	} else if numbers[middle] == number {
		return true
	}

	return false
}

func main() {
	numbers := []int{3489239843, 1, 3432, 4, 2, 34, 56565, 4, 65, 243, 432, 234}
	fmt.Println(binarySearch(numbers, 0))
}

// Дальще идёт зависимость из предыдущего задания

func partition17(numbers []int, start int, end int) int {
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

func Quicksort17(numbers []int, start int, end int) []int {
	if start > end {
		return numbers
	}
	pivot := partition17(numbers, start, end)
	numbers = Quicksort17(numbers, start, pivot-1)
	numbers = Quicksort17(numbers, pivot+1, end)
	return numbers
}
