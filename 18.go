package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
}

func increment(c *Counter, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	c.value++
	mu.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := Counter{0}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&counter, &mu, &wg)
	}

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go increment(&counter, &mu, &wg)
	}

	wg.Wait()
	fmt.Println(counter.value)

}
