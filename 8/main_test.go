package main

import (
	"testing"
	"time"
)

func TestCustomWaitGroup(t *testing.T) {
	customWait := NewCustomWaitGroup()
	blocked := make(chan struct{})

	customWait.Add(1)
	go func() {
		customWait.Wait()
		close(blocked)
	}()

	select {
	case <-blocked:
		t.Error("Expected goroutine to block on customWait.Wait()")
	default:
	}
	customWait.Done()
	for i := 0; i < 3; i++ {
		customWait.Add(1)
		go func(i int) {
			defer customWait.Done()
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	customWait.Wait()
}

func TestCustomWaitGroup_PanicOnDoneMoreThanAdd(t *testing.T) {
	customWait := NewCustomWaitGroup()
	customWait.Add(1)

	go func() {
		defer customWait.Done()
		customWait.Done()
	}()
	customWait.Wait()
}
