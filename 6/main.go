package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumbers := make(chan int)
	go GiveRandomNumbers(256, 8, randomNumbers)
	for {
		num, ok := <-randomNumbers
		if !ok {
			break
		}
		fmt.Println(num)
	}
}

func GiveRandomNumbers(rang, num int, ch chan int) chan int {
	for i := 0; i < num; i++ {
		ch <- rand.Intn(rang)
	}
	close(ch)
	return ch
}
