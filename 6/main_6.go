package main

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomNumberGenerator запускает генератор случайных чисел через небуферизированный канал
func RandomNumberGenerator(done <-chan bool) <-chan int {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	ch := make(chan int)

	go func() {
		for {
			select {
			case <-done: // Завершение работы, если получен сигнал
				close(ch)
				return
			case ch <- rand.Intn(100): // Генерация случайного числа от 0 до 99
			}
		}
	}()
	return ch
}

func main() {
	// Создаем канал завершения
	done := make(chan bool)

	// Запускаем генератор
	randGen := RandomNumberGenerator(done)

	// Получаем 5 случайных чисел
	for i := 0; i < 5; i++ {
		fmt.Println(<-randGen)
	}

	// Завершаем генератор
	done <- true
}
