package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}
	result := sliceExample(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("sliceExample() = %v, expected %v", result, expected)
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	element := 4
	expected := []int{1, 2, 3, 4}
	result := addElements(input, element)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("addElements() = %v, expected %v", result, expected)
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3}
	copied := copySlice(input)
	if !reflect.DeepEqual(copied, input) {
		t.Errorf("copySlice() = %v, expected %v", copied, input)
	}
	// Проверяем независимость копии
	input[0] = 999
	if reflect.DeepEqual(copied, input) {
		t.Errorf("Изменения в оригинальном слайсе отразились на копии")
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	index := 2
	expected := []int{1, 2, 4, 5}
	result := removeElement(input, index)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("removeElement() = %v, expected %v", result, expected)
	}
}
