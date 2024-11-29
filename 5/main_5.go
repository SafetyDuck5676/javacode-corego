package main

import "fmt"

// Функция для проверки пересечений между двумя слайсами
func findIntersection(a, b []int) (bool, []int) {
	// Создаем карту для хранения элементов из первого слайса
	elements := make(map[int]bool)
	for _, val := range a {
		elements[val] = true
	}

	// Ищем пересечения
	var intersection []int
	for _, val := range b {
		if elements[val] {
			intersection = append(intersection, val)
		}
	}

	// Возвращаем результат
	if len(intersection) == 0 {
		intersection = []int{} // Нормализуем, чтобы вернуть инициализированный пустой срез
	}

	return len(intersection) > 0, intersection
}

func main() {
	// Пример входных данных
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	// Проверяем пересечения
	hasIntersection, intersection := findIntersection(a, b)
	fmt.Println("Есть пересечения?", hasIntersection)
	fmt.Println("Пересечения:", intersection)
}
