package main

import (
	"fmt"
	"sync"
)

// mergeChannels объединяет несколько каналов в один
func mergeChannels(channels ...<-chan int) <-chan int {
	merged := make(chan int)

	var wg sync.WaitGroup

	// Функция для чтения данных из канала и записи в общий канал
	readChannel := func(ch <-chan int) {
		defer wg.Done()
		for val := range ch {
			merged <- val
		}
	}

	// Добавляем все каналы в группу ожидания
	wg.Add(len(channels))
	for _, ch := range channels {
		go readChannel(ch)
	}

	// Закрываем общий канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	// Пример работы функции mergeChannels
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Запуск горутин для наполнения каналов
	go func() {
		defer close(ch1)
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 10; i < 15; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer close(ch3)
		for i := 20; i < 25; i++ {
			ch3 <- i
		}
	}()

	// Слияние каналов
	merged := mergeChannels(ch1, ch2, ch3)

	// Вывод всех значений из объединенного канала
	for val := range merged {
		fmt.Println(val)
	}
}
