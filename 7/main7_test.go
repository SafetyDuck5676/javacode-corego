package main

import (
	"reflect"
	"sort"
	"sync"
	"testing"
)

// mergeChannels объединяет несколько каналов в один
func TestmergeChannels(channels ...<-chan int) <-chan int {
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

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Заполняем каналы
	go func() {
		defer close(ch1)
		for i := 0; i < 3; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 10; i < 13; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer close(ch3)
		for i := 20; i < 23; i++ {
			ch3 <- i
		}
	}()

	// Ожидаемый результат
	expected := []int{0, 1, 2, 10, 11, 12, 20, 21, 22}

	// Слияние каналов
	merged := mergeChannels(ch1, ch2, ch3)

	// Получение результата
	var result []int
	for val := range merged {
		result = append(result, val)
	}

	// Сортируем результат перед сравнением
	sort.Ints(result)
	sort.Ints(expected)

	// Проверяем, совпадает ли результат
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMergeEmptyChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Закрываем пустые каналы сразу после инициализации
	go func() {
		defer close(ch1)
	}()

	go func() {
		defer close(ch2)
	}()

	// Ожидаемый результат — пустой срез
	expected := []int{}

	// Слияние каналов
	merged := mergeChannels(ch1, ch2)

	// Получение результата
	var result []int
	for val := range merged {
		result = append(result, val)
	}

	// Проверяем, совпадает ли результат
	// Если сливаются пустые каналы, результат должен быть пустым срезом
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
