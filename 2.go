package main

import (
	"fmt"
	"sync"
	"time"
)

func calculatePowerOf2(x int) {
	time.Sleep(time.Second) // задержка, чтобы показать конкурентность
	fmt.Printf("%v*%v = %v\n", x, x, x*x)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, number := range numbers {
		wg.Add(1)
		go func(x int) {
			// конкуретно вызываем нашу функцию. Можно и передавать вейтгруппу как аргумент,
			// я решил показать, что можно разными способами

			calculatePowerOf2(x)
			wg.Done()
		}(number)
	}

	wg.Wait()
}
