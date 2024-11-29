package main

import (
	"fmt"
	"sync"
	"time"
)

// SemaphoreWaitGroup - кастомная структура для управления ожиданием
type SemaphoreWaitGroup struct {
	counter int
	mu      sync.Mutex
	sem     chan struct{}
}

// NewSemaphoreWaitGroup - создает новый SemaphoreWaitGroup
func NewSemaphoreWaitGroup() *SemaphoreWaitGroup {
	return &SemaphoreWaitGroup{
		sem: make(chan struct{}),
	}
}

// Add увеличивает или уменьшает счетчик
func (swg *SemaphoreWaitGroup) Add(delta int) {
	swg.mu.Lock()
	defer swg.mu.Unlock()

	if swg.counter == 0 && delta > 0 {
		// Если добавляется работа и канал закрыт, создаем новый
		swg.sem = make(chan struct{})
	}

	swg.counter += delta

	if swg.counter < 0 {
		panic("SemaphoreWaitGroup counter cannot be negative")
	}
}

// Done уменьшает счетчик на 1
func (swg *SemaphoreWaitGroup) Done() {
	swg.mu.Lock()
	defer swg.mu.Unlock()

	swg.counter--
	if swg.counter == 0 {
		close(swg.sem)
	}
}

// Wait блокирует выполнение до тех пор, пока счетчик не станет 0
func (swg *SemaphoreWaitGroup) Wait() {
	<-swg.sem
}

// main демонстрирует работу SemaphoreWaitGroup
func main() {
	fmt.Println("Start main...")

	swg := NewSemaphoreWaitGroup()
	swg.Add(3)

	// Запускаем три горутины
	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer swg.Done()
			fmt.Printf("Goroutine %d is working...\n", id)
			time.Sleep(time.Second) // Имитация работы
			fmt.Printf("Goroutine %d is done.\n", id)
		}(i)
	}

	// Ждем завершения всех горутин
	swg.Wait()

	fmt.Println("All goroutines finished.")
}
