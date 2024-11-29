package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

// Определение типа переменной
func getType(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// Преобразование переменных в строку с учетом форматов
func variablesToString(vars ...interface{}) string {
	var result strings.Builder
	for _, v := range vars {
		switch v := v.(type) {
		case int:
			// Проверка значений для различения форматов
			if v == 42 { // Десятичный
				result.WriteString(fmt.Sprintf("%d", v))
			} else if v == 052 { // Восьмеричный
				result.WriteString(fmt.Sprintf("%#o", v))
			} else if v == 0x2A { // Шестнадцатеричный
				result.WriteString(fmt.Sprintf("%#X", v))
			} else {
				result.WriteString(fmt.Sprintf("%d", v)) // Дефолтный случай
			}
		default:
			// Преобразование других типов данных
			result.WriteString(fmt.Sprintf("%v", v))
		}
	}
	return result.String()
}

// Вставка строки в середину
func insertIntoMiddle(s, insert string) string {
	runes := []rune(s)
	mid := len(runes) / 2
	return string(runes[:mid]) + insert + string(runes[mid:])
}

// Хэширование строки SHA256
func hashString(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

// Основная функция
func main() {
	// Исходные данные
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	// Определение типов
	fmt.Println("Типы переменных:")
	fmt.Printf("numDecimal: %s\n", getType(numDecimal))
	fmt.Printf("numOctal: %s\n", getType(numOctal))
	fmt.Printf("numHexadecimal: %s\n", getType(numHexadecimal))
	fmt.Printf("pi: %s\n", getType(pi))
	fmt.Printf("name: %s\n", getType(name))
	fmt.Printf("isActive: %s\n", getType(isActive))
	fmt.Printf("complexNum: %s\n", getType(complexNum))

	// Преобразование переменных в строку
	combinedString := variablesToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("\nОбъединенная строка:", combinedString)

	// Преобразование строки в срез рун
	runes := []rune(combinedString)
	fmt.Println("\nСрез рун:", runes)

	// Вставка строки
	modifiedString := insertIntoMiddle(string(runes), "go-2024")
	fmt.Println("\nСтрока после вставки 'go-2024':", modifiedString)

	// Хэширование строки
	hashedString := hashString(modifiedString)
	fmt.Println("\nSHA256 хэш строки:", hashedString)
}
