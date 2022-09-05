package main

import (
	"fmt"
	"time"
)

func write5(c chan<- string) {
	// Тут 5 на конце названия функции, чтобы иде не ругалось, что такая уже существует в пакете Мейн.
	// 5, потому что это задание номер 5

	for i := 0; true; i++ {
		time.Sleep(time.Millisecond * 2)

		item := fmt.Sprintf("item %v", i)
		fmt.Printf("\nWritten: %v\n", item) // принтим, что мы записали мменно этот айтам
		c <- item
	}
}

func read5(c <-chan string) {
	for item := range c {
		fmt.Printf("Read item: %v\n", item) // принтим, что мы прочитали тот айтем, который нам нужен
	}
}

func run5() {
	channel := make(chan string)
	go write5(channel)
	read5(channel)
}

func main() {
	done := make(chan bool, 1)
	const secondsUntilStop = 2

	go func() {
		time.Sleep(time.Second * secondsUntilStop)
		done <- true
	}()

	go run5()
	<-done
	fmt.Println("Time's up!")
}
