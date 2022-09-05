package main

import (
	"fmt"
	"sync"
)

func write9(numbers []int, ch chan<- int, wg *sync.WaitGroup) {
	wg.Add(1)
	for _, i := range numbers {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func readWrite(readCh <-chan int, writeCh chan<- int, wg *sync.WaitGroup) {
	wg.Add(1)
	for num := range readCh {
		writeCh <- num * num
	}
	close(writeCh)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan int)
	ch2 := make(chan int)
	numbers := []int{1, 2, 3}

	go write9(numbers, ch1, &wg)
	go readWrite(ch1, ch2, &wg)

	for num := range ch2 {
		fmt.Println(num)
	}
	wg.Wait()
}
