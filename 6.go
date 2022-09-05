package main

import (
	"fmt"
	"sync"
)

func returnStop(wg *sync.WaitGroup) {
	fmt.Println("Stopped goroutine by just returning")
	wg.Done()
	return
}

func chanStop(done chan bool, wg *sync.WaitGroup) {
	for value := range done {
		if value {
			break
		}
	}

	fmt.Println("Received stop signal from channel")
	wg.Done()
}

func panicStop(wg *sync.WaitGroup) {
	//panic("Panic stop")
	// Закомментировал, чтобы паников не было в рантайме

	fmt.Println("Panic stop")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	chann := make(chan bool, 1)

	go chanStop(chann, &wg)
	go returnStop(&wg)
	go panicStop(&wg)

	chann <- true
	wg.Wait()
}
