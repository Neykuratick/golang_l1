package main

import (
	"fmt"
	"sync"
	"time"
)

// Имплементация предыдущего задания, только с каналами, а не возващаемыми значениями

func calculatePower2(x int, channel chan int) {
	time.Sleep(time.Second)
	channel <- x * x
}

func calculateSum2(numbers []int) int {
	var wg sync.WaitGroup
	wg.Add(len(numbers)) // можно добавлять так, а можно и в цикле, как я делал в прошлый раз
	var sum int

	for _, number := range numbers {
		go func(x int) {
			out := make(chan int)
			go calculatePower2(x, out)
			sum += <-out
			close(out)
			wg.Done()
		}(number)
	}

	wg.Wait()
	return sum
}

func main() {
	now := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			mySlice := []int{2, 4, 6, 8, 10}
			sum := calculateSum2(mySlice)
			fmt.Printf("%v is the sum of %v\n", sum, mySlice)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(now).Seconds())
}
