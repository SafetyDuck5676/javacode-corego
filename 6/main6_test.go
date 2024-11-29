package main

import (
	"testing"
	"time"
)

func TestRandomNumberGenerator(t *testing.T) {
	done := make(chan bool)
	randGen := RandomNumberGenerator(done)

	// Проверяем получение нескольких значений
	var results []int
	for i := 0; i < 5; i++ {
		select {
		case num := <-randGen:
			results = append(results, num)
		case <-time.After(1 * time.Second):
			t.Fatal("Timeout waiting for random number")
		}
	}

	// Проверяем, что значения генерируются
	if len(results) != 5 {
		t.Errorf("Expected 5 numbers, got %d", len(results))
	}

	// Завершаем генератор
	done <- true
}

func TestGeneratorStops(t *testing.T) {
	done := make(chan bool)
	randGen := RandomNumberGenerator(done)

	// Завершаем генератор и проверяем, что канал закрывается
	done <- true

	// Ожидаем закрытия канала
	select {
	case _, ok := <-randGen:
		if ok {
			t.Error("Expected channel to be closed, but it is still open")
		}
	case <-time.After(1 * time.Second):
		t.Fatal("Timeout waiting for channel to close")
	}
}
