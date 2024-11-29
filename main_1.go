package main

import "fmt"

// StringIntMap представляет структуру данных для хранения пар "строка - число".
type StringIntMap struct {
	data map[string]int
}

// Конструктор для создания новой StringIntMap
func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

// Метод Add добавляет новую пару "ключ-значение" в карту
func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

// Метод Remove удаляет элемент по ключу из карты
func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

// Метод Copy возвращает копию карты
func (m *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int)
	for k, v := range m.data {
		newMap[k] = v
	}
	return newMap
}

// Метод Exists проверяет, существует ли ключ в карте
func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}

// Метод Get возвращает значение по ключу и булевый флаг успешности операции
func (m *StringIntMap) Get(key string) (int, bool) {
	value, exists := m.data[key]
	return value, exists
}

// Пример использования StringIntMap
func main() {
	// Создаем новый экземпляр StringIntMap
	m := NewStringIntMap()

	// Добавляем элементы
	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)
	fmt.Println("Карта после добавления элементов:", m.data)

	// Проверяем наличие ключа
	fmt.Println("Существует ли ключ 'two'? ->", m.Exists("two"))
	fmt.Println("Существует ли ключ 'four'? ->", m.Exists("four"))

	// Получаем значение по ключу
	value, found := m.Get("three")
	fmt.Printf("Значение для ключа 'three': %d (найдено: %t)\n", value, found)

	// Копируем карту
	copiedMap := m.Copy()
	fmt.Println("Копия карты:", copiedMap)

	// Удаляем элемент
	m.Remove("two")
	fmt.Println("Карта после удаления ключа 'two':", m.data)

	// Проверяем независимость копии
	fmt.Println("Копия карты после изменения оригинала:", copiedMap)
}
