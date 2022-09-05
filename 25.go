package main

import (
	"fmt"
	"sync"
	"time"
)

func sleep(d time.Duration) {
	stop := time.Now().Add(d)
	done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for {
			if time.Now().Equal(stop) {
				done <- true
				wg.Done()
			}
		}
	}()

	for value := range done {
		if value {
			break
		}
	}

	wg.Wait()
}

func main() {
	now := time.Now()
	sleep(time.Second * 2)
	fmt.Println(time.Now().Sub(now).Seconds())
}
