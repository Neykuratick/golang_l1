package main

import (
	"fmt"
	"sync"
)

func main() {
	myMap := make(map[int]string)
	var wg sync.WaitGroup
	var mu sync.Mutex // Можно юзать RWMutex, но тут это будет бесполезно, т.к тут только запись

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(index int) {
			value := fmt.Sprintf("item %v", index)

			mu.Lock()
			myMap[index] = value
			mu.Unlock()

			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(myMap)
}
