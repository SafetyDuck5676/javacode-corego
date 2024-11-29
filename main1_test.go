package main

import (
	"testing"
)

func TestToStringAll(t *testing.T) {
	expected := "4234423.14Golangtrue(1+2i)"
	result := toStringAll(42, 034, 0x2A, 3.14, "Golang", true, complex64(1+2i))
	if result != expected {
		t.Errorf("toStringAll() = %v, expected %v", result, expected)
	}
}

func TestToRunes(t *testing.T) {
	input := "hello"
	expected := []rune{'h', 'e', 'l', 'l', 'o'}
	result := toRunes(input)
	for i, r := range result {
		if r != expected[i] {
			t.Errorf("toRunes() = %v, expected %v", result, expected)
			return
		}
	}
}

func TestHashWithSalt(t *testing.T) {
	runes := []rune{'t', 'e', 's', 't'}
	salt := "salt"
	expected := "6c78d2f87a5a2e6628c9a92b6331f950b1ddf67ff0e8971e410a2b8fa1011c98"
	result := hashWithSalt(runes, salt)
	if result != expected {
		t.Errorf("hashWithSalt() = %v, expected %v", result, expected)
	}
}
