package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randomNumbers := make(chan int)
	go func() {
		for {
			num := rand.Intn(100)
			randomNumbers <- num
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		num := <-randomNumbers
		fmt.Println(num)
	}
}
