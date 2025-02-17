package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

func main() {
	num := Numbers(8)
	ch1 := make(chan uint8)
	ch2 := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(3)
	go WriteFirstChannel(&wg, ch1, num)
	go SendToSecondChannel(&wg, ch1, ch2)
	go PrintConsole(&wg, ch2)
	wg.Wait()
}

func WriteFirstChannel(wg *sync.WaitGroup, ch chan uint8, num []uint8) {
	defer wg.Done()
	for _, v := range num {
		ch <- v
	}
	close(ch)

}

func SendToSecondChannel(wg *sync.WaitGroup, ch1 chan uint8, ch2 chan float64) {
	defer wg.Done()
	for {
		n, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- CubeFloat64(n)
	}
	close(ch2)
}
func PrintConsole(wg *sync.WaitGroup, ch2 chan float64) {
	defer wg.Done()
	for res := range ch2 {
		fmt.Printf("Результат: %.2f\n", res)
	}
}
func Numbers(num int) []uint8 {
	slice := make([]uint8, num)
	for i := range slice {
		slice[i] = uint8(rand.Intn(256))
	}
	return slice
}

func CubeFloat64(num uint8) float64 {
	res := math.Pow(float64(num), 3)
	return res
}
