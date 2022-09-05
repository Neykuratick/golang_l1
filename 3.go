package main

import (
	"fmt"
	"sync"
	"time"
)

func calculatePower(x int) int {
	time.Sleep(time.Second)
	return x * x
}

func calculateSum(nums []int) int {
	var wg sync.WaitGroup
	var sum int

	for _, num := range nums {
		wg.Add(1)
		go func(x int) {
			sum += calculatePower(x)
			wg.Done()
		}(num)
	}

	wg.Wait()
	return sum
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		mySlice := []int{2, 4, 6, 8, 10}

		go func() {
			result := calculateSum(mySlice)
			fmt.Printf("%v is the sum of %v\n", result, mySlice)
			wg.Done()
		}()
	}

	wg.Wait()
}
