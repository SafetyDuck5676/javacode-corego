package main

import (
	"reflect"
	"testing"
)

func TestAddAndGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 100)
	m.Add("key2", 200)

	if val, exists := m.Get("key1"); !exists || val != 100 {
		t.Errorf("Get('key1') = %v, %v; expected 100, true", val, exists)
	}

	if val, exists := m.Get("key2"); !exists || val != 200 {
		t.Errorf("Get('key2') = %v, %v; expected 200, true", val, exists)
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 100)
	m.Remove("key1")

	if _, exists := m.Get("key1"); exists {
		t.Errorf("Remove('key1') failed; key1 should not exist")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 100)

	if !m.Exists("key1") {
		t.Errorf("Exists('key1') = false; expected true")
	}

	if m.Exists("key2") {
		t.Errorf("Exists('key2') = true; expected false")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key1", 100)
	m.Add("key2", 200)

	copied := m.Copy()

	// Проверяем, что копия идентична оригиналу
	if !reflect.DeepEqual(m.data, copied) {
		t.Errorf("Copy() = %v; expected %v", copied, m.data)
	}

	// Проверяем, что изменения в оригинале не влияют на копию
	m.Remove("key1")
	if _, exists := copied["key1"]; !exists {
		t.Errorf("Original map changed; copied map should remain unchanged")
	}
}
