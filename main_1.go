package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
)

// Преобразование переменных в строку
func toStringAll(vars ...interface{}) string {
	var result string
	for _, v := range vars {
		result += fmt.Sprintf("%v", v)
	}
	return result
}

// Преобразование строки в срез рун
func toRunes(s string) []rune {
	return []rune(s)
}

// Хэширование с добавлением соли
func hashWithSalt(runes []rune, salt string) string {
	middle := len(runes) / 2
	saltedRunes := append(runes[:middle], []rune(salt)...)
	saltedRunes = append(saltedRunes, runes[middle:]...)

	hash := sha256.Sum256([]byte(string(saltedRunes)))
	return hex.EncodeToString(hash[:])
}

// Тестовая функция
func test() {
	fmt.Println("=== Тестирование программы ===")
	var (
		numDecimal     int       = 42
		numOctal       int       = 052
		numHexadecimal int       = 0x2A
		pi             float64   = 3.14
		name           string    = "Golang"
		isActive       bool      = true
		complexNum     complex64 = 1 + 2i
	)

	vars := []interface{}{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}

	// Определение типов
	fmt.Println("Типы переменных:")
	for _, v := range vars {
		fmt.Printf("Значение: %v, Тип: %s\n", v, reflect.TypeOf(v).String())
	}

	// Преобразование в строку
	allInString := toStringAll(vars...)
	fmt.Printf("\nОбъединенная строка: %s\n", allInString)

	// Преобразование строки в срез рун
	runes := toRunes(allInString)
	fmt.Printf("\nСрез рун: %v\n", runes)

	// Хэширование с солью
	salt := "go-2024"
	hashed := hashWithSalt(runes, salt)
	fmt.Printf("\nХэш (с солью '%s'): %s\n", salt, hashed)
}

func main() {
	test()
}
