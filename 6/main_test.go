package main

import (
	"testing"
)

func TestGiveRandomNumbers(t *testing.T) {
	rang := 100
	num := 5
	ch := make(chan int)

	go GiveRandomNumbers(rang, num, ch)

	count := 0
	for range ch {
		count++
	}

	if count != num {
		t.Errorf("Expected %d random numbers, but got %d", num, count)
	}
}

func TestGiveRandomNumbersWithRangeOne(t *testing.T) {
	rang := 1
	num := 10
	ch := make(chan int)

	go GiveRandomNumbers(rang, num, ch)

	for i := 0; i < num; i++ {
		value := <-ch
		if value != 0 {
			t.Errorf("Expected 0, but got %d", value)
		}
	}

	_, ok := <-ch
	if ok {
		t.Error("Channel should be closed")
	}
}

func TestGiveRandomNumbersRange(t *testing.T) {
	rang := 100
	num := 1000
	ch := make(chan int)

	go GiveRandomNumbers(rang, num, ch)

	for n := range ch {
		if n < 0 || n >= rang {
			t.Errorf("Generated number %d is outside the range [0, %d)", n, rang)
		}
	}
}
