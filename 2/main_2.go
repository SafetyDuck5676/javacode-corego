package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для фильтрации четных чисел
func sliceExample(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

// Функция для добавления элемента в конец слайса
func addElements(slice []int, element int) []int {
	return append(slice, element)
}

// Функция для копирования слайса
func copySlice(slice []int) []int {
	copyOfSlice := make([]int, len(slice))
	copy(copyOfSlice, slice) // Используем встроенную функцию copy
	return copyOfSlice
}

// Функция для удаления элемента по индексу
func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice // Возвращаем неизменный слайс, если индекс некорректный
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	// Создаем случайный слайс
	rand.Seed(time.Now().UnixNano())
	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100) // Генерация случайных чисел от 0 до 99
	}

	fmt.Println("Оригинальный слайс:", originalSlice)

	// Фильтруем четные числа
	evenSlice := sliceExample(originalSlice)
	fmt.Println("Слайс с четными числами:", evenSlice)

	// Добавляем элемент в конец слайса
	addedSlice := addElements(originalSlice, 42)
	fmt.Println("Слайс после добавления элемента:", addedSlice)

	// Копируем слайс
	copiedSlice := copySlice(originalSlice)
	fmt.Println("Копия слайса:", copiedSlice)

	// Проверяем независимость копии
	originalSlice[0] = 999
	fmt.Println("Оригинальный слайс после изменения:", originalSlice)
	fmt.Println("Копия слайса после изменения оригинала:", copiedSlice)

	// Удаляем элемент из слайса
	indexToRemove := 3
	removedSlice := removeElement(originalSlice, indexToRemove)
	fmt.Printf("Слайс после удаления элемента с индекса %d: %v\n", indexToRemove, removedSlice)
}
