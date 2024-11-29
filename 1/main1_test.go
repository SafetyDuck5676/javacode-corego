package main

import (
	"testing"
)

func TestGetType(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{42, "int"},
		{3.14, "float64"},
		{"Golang", "string"},
		{true, "bool"},
		{complex64(1 + 2i), "complex64"},
	}
	for _, test := range tests {
		if result := getType(test.input); result != test.expected {
			t.Errorf("getType(%v) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestVariablesToString(t *testing.T) {
	result := variablesToString(42, 052, 0x2A, 3.14, "Golang", true, complex64(1+2i))
	expected := "4242420520X2A3.14Golangtrue(1+2i)" // Ожидаемая строка с разными системами счисления
	if result != expected {
		t.Errorf("variablesToString(...) = %s; expected %s", result, expected)
	}
}

func TestInsertIntoMiddle(t *testing.T) {
	result := insertIntoMiddle("abcdef", "go-2024")
	expected := "abcgo-2024def"
	if result != expected {
		t.Errorf("insertIntoMiddle(\"abcdef\", \"go-2024\") = %s; expected %s", result, expected)
	}
}

func TestHashString(t *testing.T) {
	input := "abcgo-2024def"
	expected := "08064484215e8c33b91bd482b299cb00c262161107bbb9adf980a986f8cbd718" // Убедитесь, что значение корректное
	if result := hashString(input); result != expected {
		t.Errorf("hashString(\"%s\") = %s; expected %s", input, result, expected)
	}
}
