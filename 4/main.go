package main

import (
	"fmt"
	"slices"
)

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(Difference(slice1, slice2))
}

func Difference(first, second []string) []string {
	result := make([]string, 0, len(first))
	for _, s := range first {
		if !slices.Contains(second, s) {
			result = append(result, s)
		}
	}
	return result
}
