package main

import (
	"math"
	"reflect"
	"sync"
	"testing"
)

// Тест для pipeline
func TestPipeline(t *testing.T) {
	inputChan := make(chan uint8)
	outputChan := make(chan float64)

	var wg sync.WaitGroup
	wg.Add(1)

	// Запускаем pipeline
	go pipeline(inputChan, outputChan, &wg)

	// Подаем данные
	go func() {
		input := []uint8{1, 2, 3, 4}
		for _, v := range input {
			inputChan <- v
		}
		close(inputChan)
	}()

	// Чтение результатов
	var results []float64
	go func() {
		wg.Wait()
		close(outputChan)
	}()

	for res := range outputChan {
		results = append(results, res)
	}

	// Ожидаемый результат
	expected := []float64{math.Pow(1, 3), math.Pow(2, 3), math.Pow(3, 3), math.Pow(4, 3)}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %v, got %v", expected, results)
	}
}
