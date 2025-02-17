package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		a      []int
		b      []int
		want   []int
		result bool
	}{
		{
			a:      []int{65, 3, 58, 678, 64},
			b:      []int{64, 2, 3, 43},
			want:   []int{3, 64},
			result: true,
		},
		{
			a:      []int{1, 2, 3},
			b:      []int{4, 5, 6},
			want:   []int{},
			result: false,
		},
		{
			a:      []int{10, 20, 30, 40},
			b:      []int{30, 40, 50},
			want:   []int{30, 40},
			result: true,
		},
		{
			a:      []int{7, 8, 9},
			b:      []int{1, 2, 3},
			want:   []int{},
			result: false,
		},
		{
			a:      []int{},
			b:      []int{1, 2, 3},
			want:   []int{},
			result: false,
		},
		{
			a:      []int{1, 2, 3},
			b:      []int{},
			want:   []int{},
			result: false,
		},
		{
			a:      []int{1, 2, 3, 4, 5},
			b:      []int{5, 6, 7},
			want:   []int{5},
			result: true,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			gotResult, got := Intersection(tt.a, tt.b)
			if gotResult != tt.result || !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection(%v, %v) = (%v, %v); want (%v, %v)", tt.a, tt.b, gotResult, got, tt.result, tt.want)
			}
		})
	}
}
