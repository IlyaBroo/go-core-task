package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	orSlice := originalSlice()
	chetSlice := sliceExample(orSlice)
	newslice := addElements(chetSlice, 1)
	copyslice := copySlices(newslice)
	deletedslice, err := removeElement(copyslice, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("original: ", orSlice, "\nchetnyi slice: ", chetSlice, "\nnew slice: ", newslice, "\ncopy slice", copyslice, "\ndeleted slice: ", deletedslice)

}

func originalSlice() []int {
	slice := make([]int, 10, 10)
	for i := range slice {
		slice[i] = rand.Intn(100)
	}
	return slice
}
func sliceExample(slice []int) []int {
	newSlice := make([]int, 0, 10)
	for i := range slice {
		if slice[i]%2 == 0 {
			newSlice = append(newSlice, slice[i])
		}
	}
	return newSlice
}
func addElements(slices []int, i int) []int {
	newSlice := make([]int, 0, len(slices)+1)
	// copy(newSlice, slices)
	// newSlice = append(newSlice, i)
	for _, slice := range slices {
		newSlice = append(newSlice, slice)
	}
	newSlice = append(newSlice, i)
	return newSlice
}

func copySlices(slices []int) []int {
	newslice := make([]int, 0, len(slices))
	// copy(newslice, slices) // ну или просто)
	for _, slice := range slices {
		newslice = append(newslice, slice)
	}
	return newslice
}
func removeElement(slice []int, index int) ([]int, error) {
	newSlice := make([]int, 0, len(slice))
	if index < 0 || index >= len(slice) {
		return nil, errors.New("round range")
	}
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)
	return newSlice, nil
}
