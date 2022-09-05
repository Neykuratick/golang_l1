package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func write(data string, channel chan<- string) {
	for {
		time.Sleep(time.Millisecond * 1)
		channel <- data
	}
}

func read(prefix string, channel <-chan string) {
	for {
		value, ok := <-channel
		if !ok {
			fmt.Println("Done")
			return
		}
		fmt.Println(fmt.Sprintf("%v; Message: %v", prefix, value))
	}
}

func startPolling(workerCount int, ch chan string) {
	// Обёртка для рида, чтобы можно было назначать кол-во коркеров
	for i := 0; i < workerCount; i++ {
		prefix := fmt.Sprintf("Read by worker %v", i)
		go read(prefix, ch)
	}
}

func run() {
	// функция обёртка. Просто дёргает другие функции. Сделал, чтобы декомпозировать гигансткий мейн

	channel := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)

	prefix := fmt.Sprintf("Written by worker %v", 1)
	go write(prefix, channel)
	startPolling(10, channel)

	wg.Wait()
}

func main() {
	sigs := make(chan os.Signal, 1)                      // создаём канал для чтения сигналов ОС
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // регистрируем этот канал на два системных вызова
	done := make(chan bool, 1)                           // создаём канал, в который будем записывать состояние программы

	go func() {
		sig := <-sigs    // как только программа получила один из системных вызовов, она передала его в канал
		fmt.Println()    // переход строки после ^C
		fmt.Println(sig) // выводим системный вызор
		done <- true     // передаём сообщение в канал о том, что наша прога получила системный вызов
	}()

	go run()                          // выполняем нашу прогу в фоне
	<-done                            // до тех пор, пока что-то не прочтётся из этого канала, код дальше выполняться не будет
	fmt.Println("gracefully exiting") // т.к строчка сверху - блокирующая, этот код выполнится только после прочтения из канала
}
