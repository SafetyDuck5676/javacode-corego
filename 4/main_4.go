package main

import "fmt"

// Функция, возвращающая элементы, которые есть в первом слайсе, но отсутствуют во втором
func difference(slice1, slice2 []string) []string {
	// Создаем карту для быстрого поиска
	existsInSlice2 := make(map[string]bool)
	for _, val := range slice2 {
		existsInSlice2[val] = true
	}

	// Формируем результат, исключая элементы, которые есть во втором слайсе
	var result []string
	for _, val := range slice1 {
		if !existsInSlice2[val] {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	// Пример входных данных
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	// Получаем разницу между слайсами
	result := difference(slice1, slice2)
	fmt.Println("Элементы, которые есть в первом слайсе, но отсутствуют во втором:", result)
}
