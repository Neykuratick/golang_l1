package main

import (
	"fmt"
	"sync"
	"time"
)

// Короче, мой слип просто создаёт канал (а канал это блокирующая таска)
// Из которого мы выходим, только когда время на компе == время во время старта + разница во времени

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
