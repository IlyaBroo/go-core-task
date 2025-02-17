package main

import (
	"sync"
	"testing"
	"time"
)

func TestNumbers(t *testing.T) {
	num := 10
	result := Numbers(num)

	if len(result) != num {
		t.Errorf("Expected length %d, got %d", num, len(result))
	}

	for _, v := range result {
		if v > 255 {
			t.Errorf("Value %d is out of bounds", v)
		}
	}
}

func TestCubeFloat64(t *testing.T) {
	tests := []struct {
		input    uint8
		expected float64
	}{
		{0, 0},
		{1, 1},
		{2, 8},
		{3, 27},
		{4, 64},
		{5, 125},
	}

	for _, test := range tests {
		result := CubeFloat64(test.input)
		if result != test.expected {
			t.Errorf("CubeFloat64(%d) = %f; expected %f", test.input, result, test.expected)
		}
	}
}

func TestWriteFirstChannel(t *testing.T) {
	wg := &sync.WaitGroup{}
	ch := make(chan uint8, 5)
	num := []uint8{1, 2, 3, 4, 5}

	wg.Add(1)
	go WriteFirstChannel(wg, ch, num)

	for i, expected := range num {
		select {
		case received := <-ch:
			if received != expected {
				t.Errorf("Expected %d at position %d, but got %d", expected, i, received)
			}
		case <-time.After(time.Second):
			t.Errorf("Timeout waiting for value at position %d", i)
		}
	}

	select {
	case _, ok := <-ch:
		if ok {
			t.Error("Channel was not closed")
		}
	default:
		t.Error("Channel was not closed")
	}

	wg.Wait()
}

func TestSendToSecondChannelEmptyInput(t *testing.T) {
	wg := &sync.WaitGroup{}
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	wg.Add(1)
	go SendToSecondChannel(wg, ch1, ch2)

	close(ch1) // Simulate empty input channel

	select {
	case _, ok := <-ch2:
		if ok {
			t.Error("Expected ch2 to be closed, but it was not")
		}
	case <-time.After(time.Second):
		t.Error("Timeout waiting for ch2 to close")
	}

	wg.Wait()
}
