package main

import (
	"testing"
)

func TestStringIntMap(t *testing.T) {
	m := NewStringIntMap()

	m.Add("key1", 10)
	m.Add("key2", 20)

	if value, exists := m.Get("key1"); !exists || value != 10 {
		t.Errorf("expected 10 for key1, got %d, exists: %v", value, exists)
	}
	if value, exists := m.Get("key2"); !exists || value != 20 {
		t.Errorf("expected 20 for key2, got %d, exists: %v", value, exists)
	}
	if _, exists := m.Get("key3"); exists {
		t.Errorf("expected key3 to not exist")
	}

	if !m.Exists("key1") {
		t.Error("expected key1 to exist")
	}
	if m.Exists("key3") {
		t.Error("expected key3 to not exist")
	}

	m.Remove("key1")
	if _, exists := m.Get("key1"); exists {
		t.Error("expected key1 to be removed")
	}

	copyMap := m.Copy()
	if value, exists := copyMap["key2"]; !exists || value != 20 {
		t.Errorf("expected key2 to exist in copied map with value 20, got %d, exists: %v", value, exists)
	}
	if _, exists := copyMap["key1"]; exists {
		t.Error("expected key1 to not exist in copied map")
	}
}
