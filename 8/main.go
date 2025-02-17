package main

import (
	"fmt"
	"time"
)

type CustomWaitGroup struct {
	count  int
	signal chan struct{}
}

func NewCustomWaitGroup() *CustomWaitGroup {
	str := new(CustomWaitGroup)
	str.signal = make(chan struct{})
	return str
}

func (c *CustomWaitGroup) Add(i int) {
	c.count += i
}

func (c *CustomWaitGroup) Done() {
	c.count--
	if c.count < 0 {
		c.count = 0
		return
	}
	if c.count == 0 {
		c.signal <- struct{}{}
	}
}

func (c *CustomWaitGroup) Wait() {
	<-c.signal
}

func main() {
	customWait := NewCustomWaitGroup()

	for i := 0; i < 5; i++ {
		customWait.Add(1)
		go func(i int) {
			defer customWait.Done()
			fmt.Printf("Goroutin %d started working\n", i)
			time.Sleep(time.Second)
			fmt.Printf("Goroutin %d finished working\n", i)
		}(i)
	}

	customWait.Wait()
	fmt.Println("all goroutines have finished working")
}
