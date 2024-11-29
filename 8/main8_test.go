package main

import (
	"testing"
)

// Тест SemaphoreWaitGroup с 3 задачами
func TestSemaphoreWaitGroup(t *testing.T) {
	swg := NewSemaphoreWaitGroup()
	swg.Add(3)

	go func() {
		swg.Done()
	}()
	go func() {
		swg.Done()
	}()
	go func() {
		swg.Done()
	}()

	swg.Wait()

	if swg.counter != 0 {
		t.Errorf("Expected counter to be 0, got %d", swg.counter)
	}
}

// Тест SemaphoreWaitGroup с добавлением задач
func TestSemaphoreWaitGroupWithAdd(t *testing.T) {
	swg := NewSemaphoreWaitGroup()
	swg.Add(2)

	go func() {
		swg.Done()
	}()
	go func() {
		swg.Done()
	}()

	swg.Wait()

	if swg.counter != 0 {
		t.Errorf("Expected counter to be 0, got %d", swg.counter)
	}
}

// Тест на паническую ситуацию (негативный counter)
func TestSemaphoreWaitGroupPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but got none")
		}
	}()

	swg := NewSemaphoreWaitGroup()
	swg.Add(-1) // Должно вызвать панику
}
