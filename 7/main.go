package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func mergeChannels(channels ...<-chan int) <-chan int {
	merged := make(chan int)
	var sync sync.WaitGroup
	go func() {
		for _, ch := range channels {
			sync.Add(1)
			go func(c <-chan int) {
				for val := range c {
					merged <- val
				}
				sync.Done()
			}(ch)
		}
		sync.Wait()
		close(merged)
	}()
	return merged
}

func generateNumbers(iter int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < iter; i++ {
			num := rand.Intn(100)
			ch <- num
			time.Sleep(time.Millisecond * 500)
		}
		close(ch)
	}()
	return ch
}

func main() {

	ch1 := generateNumbers(5)
	ch2 := generateNumbers(3)
	ch3 := generateNumbers(4)

	mergedChannel := mergeChannels(ch1, ch2, ch3)

	for num := range mergedChannel {
		fmt.Println(num)
	}
}
