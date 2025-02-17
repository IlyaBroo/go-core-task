package main

import (
	"reflect"
	"testing"
)

func TestOriginalSlice(t *testing.T) {
	slice := originalSlice()
	if len(slice) != 10 {
		t.Errorf("expected length 10, got %d", len(slice))
	}
	for _, v := range slice {
		if v < 0 || v > 99 {
			t.Errorf("expected values between 0 and 99, got %d", v)
		}
	}
}

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}
	result := sliceExample(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	element := 4
	expected := []int{1, 2, 3, 4}
	result := addElements(input, element)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestCopySlices(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 2, 3}
	result := copySlices(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	result[0] = 99
	if input[0] == 99 {
		t.Error("modifying the copy affected the original slice")
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	index := 2
	expected := []int{1, 2, 4, 5}
	result, err := removeElement(input, index)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	_, err = removeElement(input, 10)
	if err == nil {
		t.Error("expected an error for out-of-range index, got nil")
	}
}
