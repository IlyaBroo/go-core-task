package main

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestPrintType(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("printType panicked")
		}
	}()
	printType("test", 42)
}

func TestConvertString(t *testing.T) {
	result := convertString(42, 052, 0x2A, 3.14, "Golang", true, 1+2i)
	expected := "4242423.14Golangtrue(1+2i)"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestHashWithSalt(t *testing.T) {
	input := []rune("HelloWorld")
	salt := "go-2024"
	expectedHash := sha256.New()
	expectedHash.Write([]byte(string(append(input[:len(input)/2], append([]rune(salt), input[len(input)/2:]...)...))))
	expected := fmt.Sprintf("%x", expectedHash.Sum(nil))

	result := hashWithSalt(input, salt)
	if result != expected {
		t.Errorf("Expected hash %s but got %s", expected, result)
	}
}
