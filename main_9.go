package main

import (
	"fmt"
	"math"
	"sync"
)

// pipeline читает из inputChan, возводит в куб и записывает в outputChan
func pipeline(inputChan <-chan uint8, outputChan chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range inputChan {
		cube := math.Pow(float64(num), 3)
		outputChan <- cube
	}
}

func main() {
	// Создаем каналы
	inputChan := make(chan uint8)
	outputChan := make(chan float64)

	var wg sync.WaitGroup
	wg.Add(1)

	// Запускаем конвейер
	go pipeline(inputChan, outputChan, &wg)

	// Генерируем данные для inputChan
	go func() {
		for i := uint8(1); i <= 10; i++ {
			inputChan <- i
		}
		close(inputChan)
	}()

	// Чтение из outputChan
	go func() {
		wg.Wait() // Ждем завершения pipeline
		close(outputChan)
	}()

	// Выводим результаты
	for result := range outputChan {
		fmt.Printf("%.2f\n", result)
	}
}
