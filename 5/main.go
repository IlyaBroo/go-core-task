package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(Intersection(a, b))
}

func Intersection(a, b []int) (bool, []int) {
	result := false
	resSlise := make([]int, 0, len(a))

	for _, v := range a {
		if slices.Contains(b, v) {
			result = true
			resSlise = append(resSlise, v)
		}
	}
	return result, resSlise
}
