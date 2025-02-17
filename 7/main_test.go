package main

import (
	"testing"
)

func TestGenerateNumbers(t *testing.T) {

	ch := generateNumbers(5)
	count := 0
	for num := range ch {
		count++
		if num < 0 || num >= 100 {
			t.Errorf("Generated number %d is out of range [0, 100)", num)
		}
	}
	if count != 5 {
		t.Errorf("Expected 5 numbers, got %d", count)
	}
}

func TestMergeChannels(t *testing.T) {
	ch1 := generateNumbers(5)
	ch2 := generateNumbers(3)
	ch3 := generateNumbers(4)

	mergedChannel := mergeChannels(ch1, ch2, ch3)

	expectedCount := 5 + 3 + 4
	count := 0

	for num := range mergedChannel {
		count++
		if num < 0 || num >= 100 {
			t.Errorf("Generated number %d is out of range [0, 100)", num)
		}
	}

	if count != expectedCount {
		t.Errorf("Expected %d numbers, got %d", expectedCount, count)
	}
}
