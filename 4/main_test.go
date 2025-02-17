package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		first  []string
		second []string
		want   []string
	}{
		{
			first:  []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			second: []string{"banana", "date", "fig"},
			want:   []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			first:  []string{"a", "b", "c"},
			second: []string{"b"},
			want:   []string{"a", "c"},
		},
		{
			first:  []string{"x", "y", "z"},
			second: []string{"x", "y", "z"},
			want:   []string{},
		},
		{
			first:  []string{},
			second: []string{"a", "b"},
			want:   []string{},
		},
		{
			first:  []string{"a", "b", "c"},
			second: []string{},
			want:   []string{"a", "b", "c"},
		},
		{
			first:  []string{"apple", "banana"},
			second: []string{"banana", "banana"},
			want:   []string{"apple"},
		},
		{
			first:  []string{"apple", "banana", "cherry"},
			second: []string{"banana", "fig", "cherry"},
			want:   []string{"apple"},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Difference(tt.first, tt.second)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference(%v, %v) = %v; want %v", tt.first, tt.second, got, tt.want)
			}
		})
	}
}
